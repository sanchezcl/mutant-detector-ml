package api

import (
	"log"
	"mutantDetector/api/routes"
	"mutantDetector/config"
)

func NewServer() {
	e := routes.NewRouter()
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("Could not load configs")
	}
	e.Debug = c.AppDebug
	e.Logger.Fatal(e.Start(c.GetAddress()))
}