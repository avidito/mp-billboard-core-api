package main

import (
	"log"

	"github.com/avidito/mp-billboard-core-api/pkg"
	"github.com/avidito/mp-billboard-core-api/pkg/common/config"
	"github.com/avidito/mp-billboard-core-api/pkg/common/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := database.Init(c.DBUrl)
	pkg.InitServices(app, db)

	// Base Endpoint
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(map[string]string{
			"message": "Hello World",
		})
	})

	app.Listen(c.Port)
}
