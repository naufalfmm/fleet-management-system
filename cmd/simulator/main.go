package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/naufalfmm/fleet-management-system/config"
)

type LocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", conf.MqttBrokerHost, conf.MqttBrokerPort))
	opts.SetUsername(conf.MqttBrokerUser)
	opts.SetPassword(conf.MqttBrokerPassword)
	opts.SetClientID(fmt.Sprintf("%s-simulator", conf.MqttBrokerClientID))

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.ErrorContext(context.Background(), "connect MQTT error", "err", token.Error())
		return
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	vehicleIDs := []string{
		"B1234XYZ",
		"B5678ABC",
		"B9999TJ",
		"B7777KOR",
		"B1111JKT",
	}

	baseLat := -6.1944
	baseLon := 106.8229

	log.InfoContext(context.Background(), "sending the mocking vehicle location")

Loop:
	for {
		select {
		case <-ticker.C:
			randomVehicleID := vehicleIDs[rand.Intn(len(vehicleIDs))]

			// (Variasi sekitar ~1 km dari titik pusat)
			lat := baseLat + (rand.Float64()*0.02 - 0.01)
			lon := baseLon + (rand.Float64()*0.02 - 0.01)

			payload := LocationPayload{
				VehicleID: randomVehicleID,
				Latitude:  lat,
				Longitude: lon,
				Timestamp: time.Now().Unix(),
			}

			jsonData, err := json.Marshal(payload)
			if err != nil {
				log.ErrorContext(context.Background(), "error building json", "err", err)
				continue
			}

			topic := fmt.Sprintf("fleet/vehicle/%s/location", randomVehicleID)

			// Publish dengan QoS 1 agar menjamin data terkirim
			token := client.Publish(topic, 1, false, jsonData)
			token.Wait()

			if token.Error() != nil {
				log.ErrorContext(context.Background(), "error publish the data", "err", token.Error(), "data", payload)
				continue
			}

			log.InfoContext(context.Background(), "data is published successfully", "data", payload)

		case <-quit:
			break Loop
		}
	}

	client.Disconnect(250)
	log.InfoContext(context.Background(), "closed")
}
