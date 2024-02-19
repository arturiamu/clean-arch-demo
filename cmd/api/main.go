package main

import (
	"clean-arch-demo/config"
	"clean-arch-demo/server"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	app := server.NewApp()

	if err := app.Run(config.DefaultConf.App.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
