package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/wayzeywakeup/todo-api"
	"github.com/wayzeywakeup/todo-api/pkg/handler"
	"github.com/wayzeywakeup/todo-api/pkg/repository"
	"github.com/wayzeywakeup/todo-api/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to init config file: %v", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("http server failed to start: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}