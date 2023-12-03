package grpc

import (
	"context"
	"log"
	"net"
	"time"

	proto "github.com/BerryTracer/gps-data-service/grpc/proto"
	"github.com/BerryTracer/gps-data-service/model"
	"github.com/BerryTracer/gps-data-service/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GPSServer struct {
	GPSService service.GPSService
	proto.UnimplementedGPSServiceServer
}

func NewGPSServer(gpsService service.GPSService) *GPSServer {
	return &GPSServer{
		GPSService: gpsService,
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
	gpsDataList, err := s.GPSService.FindByUserID(ctx, req.GetDeviceId())
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
	gpsDataList, err := s.GPSService.FindByUserID(ctx, req.GetUserId())
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
	gpsData := &model.GPSData{
		DeviceID:  req.GetDeviceId(),
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
		Timestamp: time.Unix(req.GetTimestamp(), 0),
		UserID:    req.GetUserId(),
	}
	err := s.GPSService.Save(ctx, gpsData)
	if err != nil {
		log.Fatalf("failed to save gps data: %v\n", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
