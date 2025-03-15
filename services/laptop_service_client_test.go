package services_test

import (
	"context"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/anujaneja/go_grpc/sample"
	"github.com/anujaneja/go_grpc/serializer"
	"github.com/anujaneja/go_grpc/services"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"testing"
)

func TestLaptopServiceClientSearchLaptop(t *testing.T) {
	t.Parallel()
	filter := &go_grpc.Filter{
		MaxPriceUsd: 2500,
		MinCpuCores: 4,
		MinCpuGhz:   2.2,
		MinRam: &go_grpc.Memory{
			Value: 6,
			Unit:  go_grpc.Memory_GIGABYTE,
		},
	}
	store := services.NewInMemoryLaptopStore()
	expectedIDs := make(map[string]bool)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()
		switch i {
		case 0:
			laptop.PriceUsd = 2600
		case 1:
			laptop.Cpu.NumberCores = 1
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Memory = &go_grpc.Memory{
				Value: 2,
				Unit:  go_grpc.Memory_GIGABYTE,
			}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = 5.0
			laptop.Memory = &go_grpc.Memory{Value: 64, Unit: go_grpc.Memory_GIGABYTE}

			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 1800
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = 5.0
			laptop.Memory = &go_grpc.Memory{Value: 64, Unit: go_grpc.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		default:
			laptop.PriceUsd = 3000
		}
		err := store.Save(laptop)
		require.NoError(t, err)
	}
	_, serverAddress := startTestLaptopService(t, store)
	laptopClient := newTestLaptopClient(t, serverAddress)

	req := &go_grpc.SearchLaptopRequest{
		Filter: filter,
	}
	res, err := laptopClient.SearchLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	count := 0
	for {
		res, err := res.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptop().GetId())
		count++
	}
	require.Equal(t, len(expectedIDs), count)
}

func TestLaptopServiceClientCreateLaptop(t *testing.T) {
	t.Parallel()
	store := services.NewInMemoryLaptopStore()
	laptopService, serverAddress := startTestLaptopService(t, store)
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

func startTestLaptopService(t *testing.T, store services.LaptopStore) (*services.LaptopService, string) {
	laptopService := services.NewLaptopService(store)

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
