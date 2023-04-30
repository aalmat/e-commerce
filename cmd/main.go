package main

import (
	"context"
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
)

func main() {
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("APP is shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on dateabase closing: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
