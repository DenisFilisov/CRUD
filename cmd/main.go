package main

import (
	_ "CRUD/docs"
	"CRUD/pkg/handler"
	"CRUD/pkg/repository"
	"CRUD/pkg/server"
	"CRUD/pkg/service"
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"
)

// @title CRUD App API
// @version 1.0
// @description API for subscriber service

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrusChan := make(chan *logrus.Entry, 100)
	logrus.AddHook(handler.NewChannelHook(logrusChan))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initializing configuration %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Can't initialize db: %s", err.Error())
	}

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.bootstrap_servers"),
		"sasl.mechanisms":   viper.GetString("kafka.sasl_mechanisms"),
		"acks":              viper.GetString("kafka.acks")})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	producerChan := make(chan kafka.Event, 10000)

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories, kafkaProducer, producerChan)
	handlers := handler.NewHandler(services)

	go func() {
		for entry := range logrusChan {
			services.Kafka.SentMessage(entry)
		}
	}()

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
			logrus.Fatal("Can't run server:", err)
		}
	}()
	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error while shutdown process: %s", err)
	}

	if err := db.Close(); err != nil {
		logrus.Printf("error while clossing DB connection: %s", err)
	}

	close(producerChan)
	close(logrusChan)
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
