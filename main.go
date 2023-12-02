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
	db, err := database.NewGPSDatabaseConnection("mongodb://root:password@localhost:27017")
	if err != nil {
		panic(err)
	}

	defer db.Disconnect()

	mongoDBAdapter := adapter.NewMongoAdapter(db.Collection)
	gpsRepository := repository.NewMongoGPSDataRepository(mongoDBAdapter)
	gpsService := service.NewGPSService(gpsRepository)
	gpsHandler := api.NewGPSHandler(gpsService)

	app := fiber.New()

	app.Post("/gps", func(c *fiber.Ctx) error {
		return gpsHandler.SaveGPSData(api.NewFiberContext(c))
	})

	app.Listen(":3000")
}
