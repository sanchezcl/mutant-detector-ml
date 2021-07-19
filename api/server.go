package api

import (
	"mutantDetector/api/routes"
	"mutantDetector/config"
	"mutantDetector/database"
)

func NewServer() {
	e := routes.NewRouter()
	c := config.NewConfig()

	database.NewDatabaseConn()
	database.Migrate()

	e.Debug = c.AppDebug
	e.Logger.Fatal(e.Start(c.GetAddress()))
}