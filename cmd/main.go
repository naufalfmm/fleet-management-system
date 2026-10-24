package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naufalfmm/fleet-management-system/config"
	routestopdomain "github.com/naufalfmm/fleet-management-system/internal/domain/routestop"
	vehicleassignmentdomain "github.com/naufalfmm/fleet-management-system/internal/domain/vehicleassignment"
	vehiclelocationdomain "github.com/naufalfmm/fleet-management-system/internal/domain/vehiclelocation"
	vehiclelocationmqttctrl "github.com/naufalfmm/fleet-management-system/internal/infrastructure/mqtt/controller/vehiclelocation"
	"github.com/naufalfmm/fleet-management-system/internal/infrastructure/mqtt/subscriber"
	middlewarerest "github.com/naufalfmm/fleet-management-system/internal/infrastructure/rest/controller/middleware"
	"github.com/naufalfmm/fleet-management-system/internal/infrastructure/rest/controller/router"
	vehiclelocationrestctrl "github.com/naufalfmm/fleet-management-system/internal/infrastructure/rest/controller/vehiclelocation"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/messagequeue/rabbitmq"
	v10 "github.com/naufalfmm/fleet-management-system/internal/pkg/validator/v10"
	routestoprepo "github.com/naufalfmm/fleet-management-system/internal/repository/routestop"
	vehicleassignmentrepo "github.com/naufalfmm/fleet-management-system/internal/repository/vehicleassignment"
	vehiclelocationrepo "github.com/naufalfmm/fleet-management-system/internal/repository/vehiclelocation"
	vehiclelocationusc "github.com/naufalfmm/fleet-management-system/internal/usecase/vehiclelocation"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ge := gin.New()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", conf.MqttBrokerHost, conf.MqttBrokerPort))
	opts.SetUsername(conf.MqttBrokerUser)
	opts.SetPassword(conf.MqttBrokerPassword)
	opts.SetClientID(conf.MqttBrokerClientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, conf.GetDSN())
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}

	mq, err := rabbitmq.NewRabbitMQ(conf.GetAmqpUrl())
	if err != nil {
		panic(err)
	}

	vld, err := v10.NewV10()
	if err != nil {
		panic(err)
	}

	routeStopQuerier := routestopdomain.New(dbpool)
	vehicleAssignmentQuerier := vehicleassignmentdomain.New(dbpool)
	vehicleLocationQuerier := vehiclelocationdomain.New(dbpool)

	routeStopRepository := routestoprepo.NewRepository(routeStopQuerier)
	vehicleAssignmentRepository := vehicleassignmentrepo.NewRepository(vehicleAssignmentQuerier)
	vehicleLocationRepository := vehiclelocationrepo.NewRepository(vehicleLocationQuerier)

	vehicleLocationUsecase := vehiclelocationusc.NewUsecase(vehicleLocationRepository, vehicleAssignmentRepository, routeStopRepository, mq, *conf)

	vehicleLocationMqttController := vehiclelocationmqttctrl.NewController(vehicleLocationUsecase, vld)

	subscriber := subscriber.NewSubscriber(vehicleLocationMqttController, *conf)
	subscriber.Register(client)

	vehicleLocationRestController := vehiclelocationrestctrl.NewController(vehicleLocationUsecase)
	restMiddleware := middlewarerest.NewMiddleware()

	router := router.NewRouter(vehicleLocationRestController, restMiddleware)
	router.Register(ge)

	binding.Validator = vld

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: ge,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sc

	client.Disconnect(250)

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctxShutDown); err != nil {
		panic(err)
	}
}
