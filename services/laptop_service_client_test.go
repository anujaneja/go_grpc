package services_test

import (
	"context"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/anujaneja/go_grpc/sample"
	"github.com/anujaneja/go_grpc/serializer"
	"github.com/anujaneja/go_grpc/services"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func TestLaptopServiceClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopService, serverAddress := startTestLaptopService(t)
	laptopClient := newTestLaptopClient(t, serverAddress)
	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &go_grpc.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.GetId())

	other, err := laptopService.Store.Find(res.GetId())
	require.NoError(t, err)
	require.NotNil(t, other)

	//check for same output...which is saved
	checkForSameLaptop(t, laptop, other)
}

func checkForSameLaptop(t *testing.T, laptop1 *go_grpc.Laptop, laptop2 *go_grpc.Laptop) {
	//require.Equal(t, laptop1, laptop2)
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)
	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)
	require.Equal(t, json1, json2)
}

func newTestLaptopClient(t *testing.T, address string) go_grpc.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	require.NoError(t, err)
	return go_grpc.NewLaptopServiceClient(conn)
}

func startTestLaptopService(t *testing.T) (*services.LaptopService, string) {
	laptopService := services.NewLaptopService(services.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	go_grpc.RegisterLaptopServiceServer(grpcServer, laptopService)
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return laptopService, listener.Addr().String()
}
