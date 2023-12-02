package main

import (
	"common-utils/adapter"
	"gps-data-receiver/api"
	"gps-data-receiver/database"
	"gps-data-receiver/repository"
	"gps-data-receiver/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new database connection
	db, err := database.NewGPSDatabaseConnection("mongodb://root:password@localhost:27017")
	if err != nil {
		panic(err)
	}
	defer db.Disconnect()

	// Create adapters, repositories, services, and handlers
	mongoDBAdapter := adapter.NewMongoAdapter(db.Collection)
	gpsRepository := repository.NewMongoGPSDataRepository(mongoDBAdapter)
	gpsService := service.NewGPSService(gpsRepository)
	gpsHandler := api.NewGPSHandler(gpsService)

	// Create a new Fiber app and router
	app := fiber.New()
	fiberRouter := api.NewFiberRouter(app)

	// Define routes
	fiberRouter.Post("/gps", gpsHandler.SaveGPSData)

	// Start the server
	app.Listen(":3000")
}
