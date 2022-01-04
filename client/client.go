package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-go/proto"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCarManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newCars []pb.NewCarParams
	newCars = append(newCars, pb.NewCarParams{})
	newCars[0].Mark = "Toyota"
	newCars[0].Model = "Corolla"
	newCars[0].Price = 2222

	for _, v := range newCars {
		r, err := c.CreateNewCar(ctx, &pb.NewCarParams{Mark: v.Mark, Model: v.Model, Price: v.Price})
		if err != nil {
			log.Fatalf("Could not create user: %v", err)
		}
		log.Printf(`User Details:
		MARK: %s
		MODEL: %s
		PRICE: %d
		ID: %d`, r.GetMark(), r.GetModel(), r.GetPrice(), r.GetId())
	}

	// GetCars
	params := &pb.GetCarsParams{}
	r, err := c.GetCars(ctx, params)
	if err != nil {
		log.Fatalf("Could not retrieve cars: %v", err)
	}
	log.Print("\nCAR LIST: \n")
	fmt.Printf("r.GetCars(): %v\n", r.GetCars())

	//GetOne

	paramsOne := &pb.GetOneCarParams{}
	paramsOne.Id = "057fd030-6e96-4dc7-9469-3c9cef364a6e"
	result, err := c.GetOneCar(ctx, paramsOne)
	if err != nil {
		log.Fatalf("Could not retrieve cars: %v", err)
	}
	log.Print("\nCAR LIST: \n")
	fmt.Printf("Result One(): %v\n", result)
}
