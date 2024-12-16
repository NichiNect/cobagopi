package main

import (
	"cobagopi/apps/auth"
	"cobagopi/external/database"
	"cobagopi/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
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

	// ? Register route auth
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
