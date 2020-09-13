package main

import (
	"github.com/rafa-acioly/challenges/api"
	"github.com/rafa-acioly/challenges/app"
	"github.com/rafa-acioly/challenges/config"
)

func main() {
	config := config.New()

	app := app.App{}
	app.Initialize(config)

	walkResource := api.NewDogWalkResource(app.DB)
	app.AddRoute("/walks", walkResource.Routes)

	app.Run(":8080")
}
