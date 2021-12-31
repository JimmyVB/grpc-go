package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	pb "grpc-go/proto"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
)

const (
	port = ":50051"
)

func NewCarManagementServer() *CarManagementServer {
	return &CarManagementServer{}
}

type CarManagementServer struct {
	pb.UnimplementedCarManagementServer
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

func (server *CarManagementServer) CreateNewCar(ctx context.Context, in *pb.NewCar) (*pb.Car, error) {
	log.Printf("Received: %v", in.GetMark())
	readBytes, err := ioutil.ReadFile("users.json")
	var carsList *pb.CarsList = &pb.CarsList{}
	var carId int32 = int32(rand.Intn(1000))
	createdCar := &pb.Car{Mark: in.GetMark(), Model: in.GetModel(), Price: in.GetPrice(), Id: carId}

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found. Creating a new file")
			carsList.Cars = append(carsList.Cars, createdCar)
			jsonBytes, err := protojson.Marshal(carsList)
			if err != nil {
				log.Fatalf("JSON Marshaling failed: %v", err)
			}
			if err := ioutil.WriteFile("cars.json", jsonBytes, 0664); err != nil {
				log.Fatalf("Failed write to file: %v", err)
			}
			return createdCar, err
		} else {
			log.Fatalln("Error reading file: ", err)
		}
	}

	if err := protojson.Unmarshal(readBytes, carsList); err != nil {
		log.Fatalf("Failed to parse car list: %v", err)
	}
	carsList.Cars = append(carsList.Cars, createdCar)
	jsonBytes, err := protojson.Marshal(carsList)
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %v", err)
	}
	if err := ioutil.WriteFile("cars.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Failed write to file: %v", err)
	}
	return createdCar, nil
}

func (s *CarManagementServer) GetCars(ctx context.Context, in *pb.GetCarsParams) (*pb.CarsList, error) {
	jsonBytes, err := ioutil.ReadFile("cars.json")
	if err != nil {
		log.Fatalf("Failed read from file: %v", err)
	}
	var carsList *pb.CarsList = &pb.CarsList{}
	if err := protojson.Unmarshal(jsonBytes, carsList); err != nil {
		log.Fatalf("Unmarshaling failed: %v", err)
	}

	return carsList, nil
}

func main() {
	var carServer *CarManagementServer = NewCarManagementServer()
	if err := carServer.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
