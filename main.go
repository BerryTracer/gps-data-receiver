package main

import (
	"log"
	"net"

	authservice "github.com/BerryTracer/auth-service/grpc/proto"
	"github.com/BerryTracer/common-service/adapter/database/mongodb"
	"github.com/BerryTracer/common-service/config"
	"github.com/BerryTracer/gps-data-service/api"
	"github.com/BerryTracer/gps-data-service/database"
	proto "github.com/BerryTracer/gps-data-service/grpc/proto"
	server "github.com/BerryTracer/gps-data-service/grpc/server"
	"github.com/BerryTracer/gps-data-service/repository"
	"github.com/BerryTracer/gps-data-service/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load environment variables from .env file
	envLoader := config.NewRealEnvLoader()

	// Load the environment variable for the MongoDB URI
	mongodbURI, err := config.LoadEnv(envLoader, "MONGODB_URI")
	if err != nil {
		panic(err)
	}

	// Create a new database connection
	db, err := database.NewGPSDatabaseConnection(mongodbURI)
	if err != nil {
		panic(err)
	}
	defer func(db *database.GPSDatabase) {
		err := db.Disconnect()
		if err != nil {
			log.Fatalf("failed to serve gRPC server: %v\n", err)
		}
	}(db)

	// Create adapters, repositories, services, and handlers
	mongoDBAdapter := mongodb.NewMongoAdapter(db.Collection)
	gpsRepository := repository.NewMongoGPSDataRepository(mongoDBAdapter)
	gpsService := service.NewGPSService(gpsRepository)
	gpsHandler := api.NewGPSHandler(gpsService)

	grpcPort, err := config.LoadEnvWithDefault(envLoader, "GRPC_PORT", "50051")
	if err != nil {
		panic(err)
	}

	// Load the environment variable for the AuthService URI
	authAuthServiceURI, err := config.LoadEnv(envLoader, "AUTH_SERVICE_URI")
	if err != nil {
		log.Fatalf("failed to load environment variable: %v", err)
	}

	// Establish a connection to the gRPC server
	conn, err := grpc.Dial(authAuthServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to create indexes: %v", err)
		}
	}(conn)

	// Create a client for the AuthService
	authServiceClient := authservice.NewAuthServiceClient(conn)

	gpsGRPCServer := server.NewGPSServer(authServiceClient, gpsService)

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

	httpPort, err := config.LoadEnvWithDefault(envLoader, "HTTP_PORT", "3000")
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
	err = app.Listen(":" + httpPort)
	if err != nil {
		log.Fatalf("failed to serve gRPC server: %v\n", err)
	}
}
