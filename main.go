package main

import (
	"github.com/BerryTracer/common-service/adapter"
	"github.com/BerryTracer/common-service/config"
	"github.com/BerryTracer/gps-data-service/api"
	"github.com/BerryTracer/gps-data-service/database"
	"github.com/BerryTracer/gps-data-service/repository"
	"github.com/BerryTracer/gps-data-service/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	mongodbURI, err := config.LoadEnv("MONGODB_URI")
	if err != nil {
		panic(err)
	}

	// Create a new database connection
	db, err := database.NewGPSDatabaseConnection(mongodbURI)
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
	fiberRouter.Get("/gps", gpsHandler.GetGPSDataByDeviceId)

	// Start the server
	app.Listen(":3000")
}
