package main

import (
	_ "github.com/lib/pq"

	"crud-app/internal/config"
	"crud-app/internal/server"
)

func main() {
	cfg, dbCfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg, dbCfg)
}
