package main

import (
	"cobagopi/external/database"
	"cobagopi/internal/config"
	"fmt"
)

func main() {
	// ? Load config
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	// ? Load DB
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		fmt.Println("Database connected.")
	}
}
