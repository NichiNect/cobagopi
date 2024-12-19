package main

import (
	"cobagopi/apps/auth"
	"cobagopi/apps/product"
	"cobagopi/apps/transaction"
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

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	// ? Register route
	auth.Init(router, db)
	product.Init(router, db)
	transaction.Init(router, db)

	// ? Listen app
	router.Listen(config.Cfg.App.Port)
}
