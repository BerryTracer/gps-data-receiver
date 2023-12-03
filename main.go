package main

import (
	"log"
	"net"

	"github.com/BerryTracer/common-service/adapter"
	"github.com/BerryTracer/common-service/config"
	"github.com/BerryTracer/gps-data-service/api"
	"github.com/BerryTracer/gps-data-service/database"
	proto "github.com/BerryTracer/gps-data-service/grpc/proto"
	server "github.com/BerryTracer/gps-data-service/grpc/server"
	"github.com/BerryTracer/gps-data-service/repository"
	"github.com/BerryTracer/gps-data-service/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
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

	grpcPort, err := config.LoadEnvWithDefault("GRPC_PORT", "50051")
	if err != nil {
		panic(err)
	}

	gpsGRPCServer := server.NewGPSServer(gpsService)

	// Run the gRPC server in a separate goroutine to not block the main thread
	go func() {
		// Listen for gRPC requests on a different port (e.g., :50051)
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v\n", err)
		}
		grpcServer := grpc.NewServer()
		proto.RegisterGPSServiceServer(grpcServer, gpsGRPCServer)

		log.Println("gRPC server listening on port " + grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC server: %v\n", err)
		}
	}()

	httpPort, err := config.LoadEnvWithDefault("HTTP_PORT", "3000")
	if err != nil {
		panic(err)
	}

	// Create a new Fiber app and router
	app := fiber.New()
	fiberRouter := api.NewFiberRouter(app)

	// Define routes
	fiberRouter.Post("/gps", gpsHandler.SaveGPSData)
	fiberRouter.Get("/gps/device/:device_id", gpsHandler.GetGPSDataByDeviceId)
	fiberRouter.Get("/gps/user/:user_id", gpsHandler.GetGPSDataByUserId)

	// Start the server
	app.Listen(":" + httpPort)
}
