package main

import (
	"github.com/Savitar465/task-manager/config"
	"github.com/Savitar465/task-manager/server"
	"github.com/spf13/viper"
	"github.com/swaggest/openapi-go/openapi3"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	app := server.NewApp()
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}

	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("Things API").
		WithVersion("1.2.3").
		WithDescription("Put something here")
}
