package main

import (
	"context"
	"flag"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/anujaneja/go_grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	laptopClient := go_grpc.NewLaptopServiceClient(conn)
	laptop := sample.NewLaptop()
	//check with for already exists laptop...
	//laptop.Id = "55ac1dde-1e06-47d5-aaf2-e48738378cc6"
	//check with invalid uuid...
	//laptop.Id = "invalid-id"
	req := go_grpc.CreateLaptopRequest{
		Laptop: laptop,
	}

	//setting timeout in request at the client level...
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := laptopClient.CreateLaptop(ctx, &req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatalf("failed to create laptop: %v", err)
		}
		return
	}
	log.Printf("laptop created with id: %v", res.GetId())
}
