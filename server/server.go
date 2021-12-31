package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	pb "grpc-go/proto"
	"log"
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
	conn *pgx.Conn
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
	createSql := `
	create table if not exists cars (
		id text PRIMARY KEY,
		marca text,
		model text,
		price int);
		`
	id := uuid.New()

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}
	createdCar := &pb.Car{Id: id.String(), Mark: in.GetMark(), Model: in.GetModel(), Price: in.GetPrice()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(),
		"insert into cars (id, marca, model, price) values ($1, $2, $3, $4)",
		createdCar.Id, createdCar.Mark, createdCar.Model, createdCar.Price)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return createdCar, nil
}

func (server *CarManagementServer) GetCars(ctx context.Context, in *pb.GetCarsParams) (*pb.CarsList, error) {

	var carsList *pb.CarsList = &pb.CarsList{}

	rows, err := server.conn.Query(context.Background(), "select * from cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		car := pb.Car{}
		err = rows.Scan(&car.Id, &car.Mark, &car.Model, &car.Price)
		if err != nil {
			return nil, err
		}
		carsList.Cars = append(carsList.Cars, &car)
	}
	return carsList, nil
}

func main() {
	database_url := "postgres://admin:admin@localhost:5432/crudgo?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())

	var carServer *CarManagementServer = NewCarManagementServer()
	carServer.conn = conn
	if err := carServer.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
