package main

import (
	"github/sgo-chat/internals/configs"
	"github/sgo-chat/internals/database"
	"github/sgo-chat/modules"
	"log"
)

func main() {
	cfg := configs.Load()

	client, err := database.ConnectMongo(cfg.MongoURL)
	if err != nil {
		log.Fatalln("Connect to Mongo Failed: " + err.Error())
	}

	collections := map[string][]string{
		"accounts": {"email"},
		"profiles": {"accountId"},
	}

	err = database.SetupUniqueIndex(client, cfg.DBname, collections)
	if err != nil {
		log.Fatalln("Setup unique index failed: " + err.Error())
	}

	db := client.Database(cfg.DBname)
	r := modules.Setup(db, cfg)

	r.Run(":" + cfg.Port)
}
