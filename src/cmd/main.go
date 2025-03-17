package main

import (
	"log"

	"github.com/alielmi98/go-ecommerce-api/api"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/data/cache"
	"github.com/alielmi98/go-ecommerce-api/data/db"
	"github.com/alielmi98/go-ecommerce-api/data/db/migrations"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatalf("caller:%s  Level:%s  Msg:%s", constants.Redis, constants.Startup, err.Error())
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatalf("caller:%s  Level:%s  Msg:%s", constants.Postgres, constants.Startup, err.Error())
	}

	migrations.Up_1()
	api.InitServer(cfg)

}
