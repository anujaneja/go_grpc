package main

import (
	"context"
	"flag"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/anujaneja/go_grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
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
	laptopClient := go_grpc.NewLaptopServiceClient(conn) //check with for already exists laptop...
	//laptop.Id = "55ac1dde-1e06-47d5-aaf2-e48738378cc6"
	//check with invalid uuid...
	//laptop.Id = "invalid-id"
	CreateLaptop(laptopClient)
	filter := &go_grpc.Filter{
		MaxPriceUsd: 5000,
		MinCpuCores: 1,
		MinCpuGhz:   1,
		MinRam:      &go_grpc.Memory{Value: 1, Unit: go_grpc.Memory_GIGABYTE},
	}
	SearchLaptop(laptopClient, filter)
}

func SearchLaptop(client go_grpc.LaptopServiceClient, filter *go_grpc.Filter) {
	log.Printf("search laptop with filter: %v", filter)
	req := &go_grpc.SearchLaptopRequest{
		Filter: filter,
	}

	//setting timeout in request at the client level...
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	stream, err := client.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatalf("failed to search laptop: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		laptop := res.GetLaptop()
		log.Printf("_ found: %v", laptop.GetId())
		log.Printf(" + brand: %v", laptop.GetBrand())
		log.Printf(" + name: %v", laptop.GetName())
		log.Printf(" + cpu cores: %v", laptop.GetCpu().GetNumberCores())
	}
}

func CreateLaptop(laptopClient go_grpc.LaptopServiceClient) {
	laptop := sample.NewLaptop()
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
