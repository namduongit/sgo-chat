package main

import (
	"github/sgo-chat/internal/config"
	"github/sgo-chat/internal/database"
	controllers "github/sgo-chat/modules"
	"log"
)

func main() {
	cfg := config.Load()

	client, err := database.ConnectMongo(cfg.MongoURL)
	if err != nil {
		log.Fatal("Connect to Mongo Failed: " + err.Error())
	}

	db := client.Database(cfg.DBname)
	r := controllers.Setup(db, cfg)

	r.Run(":" + cfg.Port)
}
