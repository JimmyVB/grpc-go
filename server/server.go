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
	port = ":10051"
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

func (server *CarManagementServer) CreateNewCar(ctx context.Context, in *pb.NewCarParams) (*pb.Car, error) {
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
		log.Fatalf("conn.Begin Created failed: %v", err)
	}
	_, err = tx.Exec(context.Background(),
		"insert into cars (id, marca, model, price) values ($1, $2, $3, $4)",
		createdCar.Id, createdCar.Mark, createdCar.Model, createdCar.Price)
	if err != nil {
		log.Fatalf("tx.Exec Created failed: %v", err)
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

func (server *CarManagementServer) GetOneCar(ctx context.Context, in *pb.GetOneCarParams) (*pb.Car, error) {
	car := &pb.Car{}

	err := server.conn.QueryRow(context.Background(),
		"select * from cars where id = $1", in.GetId()).Scan(&car.Id, &car.Mark, &car.Model, &car.Price)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (server *CarManagementServer) UpdateCar(ctx context.Context, in *pb.UpdateCarParams) (*pb.Car, error) {
	//car := &pb.Car{}
	//
	//_, err := server.conn.Exec(context.Background(),
	//	"update cars set marca = $2, model = $3, price = $4 where id = $1",
	//	in.GetMark(), in.GetModel(), in.GetPrice(), in.GetId())
	//if err != nil {
	//	return nil, err
	//}

	updateCar := &pb.Car{Id: in.GetId(), Mark: in.GetMark(), Model: in.GetModel(), Price: in.GetPrice()}

	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin Updated failed: %v", err)
	}
	_, err = tx.Exec(context.Background(),
		"update cars set marca = $2, model = $3, price = $4 where id = $1",
		updateCar.Id, updateCar.Mark, updateCar.Model, updateCar.Price)
	if err != nil {
		log.Fatalf("tx.Exec Udpated failed: %v", err)
	}
	tx.Commit(context.Background())

	return updateCar, nil
}

func main() {
	databaseUrl := "postgres://admin:admin@localhost:5432/crudgo?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
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
