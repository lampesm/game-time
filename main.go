package main

import (
	"github.com/lampesm/game-time/config"
	"github.com/lampesm/game-time/db"
	"github.com/lampesm/game-time/router"
)

func main() {
	config.LoadEnv()
	db.Connection()

	r := router.SetupRoute()
	r.Run()
}
