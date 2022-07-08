package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Shreeyash-Naik/Valet-Parking/api"
	"github.com/Shreeyash-Naik/Valet-Parking/env"
	"github.com/gofiber/fiber/v2"
)

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}

func main() {
	// Set Env variables
	env.SetEnvs()

	// Connect Database
	api.ConnectDB()

	// Migrate relations
	toMigrate, _ := strconv.ParseBool(os.Getenv("TO_MIGRATE"))
	if toMigrate {
		api.Migrate()
	}

	// Init Fiber App
	app := fiber.New()

	// Health Checks
	app.Post("/", healthCheck)
	app.Get("/health", healthCheck)

	// Mount routes
	api.MountRoutes(app)

	// Fiber Start
	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalln("Error Starting Fiber: ", err)
	}
}
