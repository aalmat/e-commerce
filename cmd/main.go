package main

import (
	"context"
	"flag"
	"fmt"
	e_commerce "github.com/aalmat/e-commerce"
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/handler"
	"github.com/aalmat/e-commerce/pkg/repository"
	"github.com/aalmat/e-commerce/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const tickInterval = time.Minute

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error init configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("DB connection error: %s", err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Product{}, &models.Rating{}, &models.WareHouse{}, &models.Commentary{}, &models.Order{})

	repos := repository.NewRepostitory(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(e_commerce.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.Routes()); err != nil {
			logrus.Fatalf("server error: %s", err.Error())
		}
	}()

	go repos.Admin.CheckOrders(tickInterval)

	// Graceful Shutdown
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT) // CTRL+C
	<-quit
	fmt.Println("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Server was successful shutdown.")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
