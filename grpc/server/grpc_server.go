package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	authservice "github.com/BerryTracer/auth-service/grpc/proto"
	proto "github.com/BerryTracer/gps-data-service/grpc/proto"
	"github.com/BerryTracer/gps-data-service/model"
	"github.com/BerryTracer/gps-data-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GPSServer struct {
	AuthService authservice.AuthServiceClient
	GPSService  service.GPSService
	proto.UnimplementedGPSServiceServer
}

func NewGPSServer(authService authservice.AuthServiceClient, gpsService service.GPSService) *GPSServer {
	return &GPSServer{
		AuthService: authService,
		GPSService:  gpsService,
	}
}

func (s *GPSServer) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return err
	}
	server := grpc.NewServer()
	proto.RegisterGPSServiceServer(server, s)
	log.Printf("gRPC server listening on port %s\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		return err
	}
	return nil
}

func (s *GPSServer) FindByDeviceID(ctx context.Context, req *proto.FindByDeviceIDRequest) (*proto.GPSDataList, error) {
	err := s.extractAndValidateToken(ctx, map[string]string{"device_id": req.GetDeviceId()})
	if err != nil {
		return nil, err
	}

	gpsDataList, err := s.GPSService.FindByUserID(ctx, req.GetDeviceId(), int64(req.Limit), int64(req.Offset))
	if err != nil {
		log.Fatalf("failed to find gps data by device id: %v\n", err)
		return nil, err
	}

	protoGPSDataList := &proto.GPSDataList{}
	for _, gpsData := range gpsDataList {
		protoGPSDataList.GpsData = append(protoGPSDataList.GpsData, gpsData.ConvertToProto())
	}

	return protoGPSDataList, nil
}

func (s *GPSServer) FindByUserID(ctx context.Context, req *proto.FindByUserIDRequest) (*proto.GPSDataList, error) {
	err := s.extractAndValidateToken(ctx, map[string]string{"user_id": req.GetUserId()})
	if err != nil {
		return nil, err
	}

	gpsDataList, err := s.GPSService.FindByUserID(ctx, req.GetUserId(), int64(req.Limit), int64(req.Offset))
	if err != nil {
		log.Fatalf("failed to find gps data by user id: %v\n", err)
		return nil, err
	}

	protoGPSDataList := &proto.GPSDataList{}
	for _, gpsData := range gpsDataList {
		protoGPSDataList.GpsData = append(protoGPSDataList.GpsData, gpsData.ConvertToProto())
	}

	return protoGPSDataList, nil
}

func (s *GPSServer) Save(ctx context.Context, req *proto.GPSData) (*emptypb.Empty, error) {
	err := s.extractAndValidateToken(ctx, map[string]string{"device_id": req.GetDeviceId(), "user_id": req.GetUserId()})
	if err != nil {
		return nil, err
	}

	gpsData := &model.GPSData{
		DeviceID:  req.GetDeviceId(),
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
		Timestamp: time.Unix(req.GetTimestamp(), 0),
		UserID:    req.GetUserId(),
	}

	if err := gpsData.Validate(); err != nil {
		log.Fatalf("failed to validate gps data: %v\n", err)
		return nil, err
	}

	err = s.GPSService.Save(ctx, gpsData)
	if err != nil {
		log.Fatalf("failed to save gps data: %v\n", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *GPSServer) extractAndValidateToken(ctx context.Context, expectedClaims map[string]string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("missing metadata from context")
	}

	if len(md["authorization"]) == 0 {
		return errors.New("missing authorization token")
	}

	token := md["authorization"][0]
	tokenResults, err := s.AuthService.VerifyToken(ctx, &authservice.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		return err
	}

	if !tokenResults.GetValid() {
		return fmt.Errorf("invalid token: %s", token)
	}

	tokenClaims := tokenResults.GetClaims()
	for key, expectedValue := range expectedClaims {
		if value, ok := tokenClaims[key]; !ok || value != expectedValue {
			return fmt.Errorf("invalid or missing %s in token", key)
		}
	}

	return nil
}
