package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-go/proto"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

func NewCarManagementServer() *CarManagementServer {
	return &CarManagementServer{
		carList: &pb.CarList{},
	}
}

type CarManagementServer struct {
	pb.UnimplementedCarManagementServer
	carList *pb.CarList
}

func (server *CarManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCarManagementServer(s, server)
	log.Printf("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *CarManagementServer) CreateNewCar(ctx context.Context, in *pb.NewCar) (*pb.Car, error) {
	log.Printf("Received: %v", in.GetMark())
	var carId int32 = int32(rand.Intn(1000))
	createdCar := &pb.Car{Mark: in.GetMark(), Model: in.GetModel(), Price: in.GetPrice(), Id: carId}
	s.carList.Cars = append(s.carList.Cars, createdCar)
	return createdCar, nil
}

func (s *CarManagementServer) GetCars(ctx context.Context, in *pb.GetCarsParams) (*pb.CarList, error) {
	return s.carList, nil
}

func main() {
	var carServer *CarManagementServer = NewCarManagementServer()
	if err := carServer.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
