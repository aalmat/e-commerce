package main

import (
	e_commerce "github.com/aalmat/e-commerce"
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/handler"
	"github.com/aalmat/e-commerce/pkg/repository"
	"github.com/aalmat/e-commerce/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
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

	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Product{})

	repos := repository.NewRepostitory(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(e_commerce.Server)
	if err := server.Run(viper.GetString("port"), handlers.Routes()); err != nil {
		logrus.Fatalf("server error: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
