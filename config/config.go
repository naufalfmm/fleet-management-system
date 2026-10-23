package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Name string `envconfig:"NAME"`
	Port int    `envconfig:"PORT" default:"8080"`

	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBUsername string `envconfig:"DB_USERNAME"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
	DBSSL      string `envconfig:"DB_SSL" default:"disable"`

	MqttBrokerHost     string `envconfig:"MQTT_BROKER_HOST" default:"localhost"`
	MqttBrokerPort     string `envconfig:"MQTT_BROKER_PORT" default:"1883"`
	MqttBrokerUser     string `envconfig:"MQTT_BROKER_USER"`
	MqttBrokerPassword string `envconfig:"MQTT_BROKER_PASSWORD"`

	MessageQueueHost     string `envconfig:"MESSAGE_QUEUE_HOST" default:"localhost"`
	MessageQueuePort     string `envconfig:"MESSAGE_QUEUE_PORT" default:"5672"`
	MessageQueueUser     string `envconfig:"MESSAGE_QUEUE_USER"`
	MessageQueuePassword string `envconfig:"MESSAGE_QUEUE_PASSWORD"`

	RabbitMqGeofencingExchange string `envconfig:"RABBITMQ_GEOFENCING_EXCHANGE"`
	RabbitMqGeofencingQueue    string `envconfig:"RABBITMQ_GEOFENCING_QUEUE"`

	GeofencingRadiusMeter float64 `envconfig:"GEOFENCING_RADIUS_METER" default:"50"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUsername,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSL,
	)
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("./config/.env")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
