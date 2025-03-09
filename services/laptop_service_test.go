package services_test

import (
	"context"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/anujaneja/go_grpc/sample"
	"github.com/anujaneja/go_grpc/services"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServiceCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopWithoutId := sample.NewLaptop()
	laptopWithoutId.Id = ""
	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid-uuid"
	laptopDuplicateId := sample.NewLaptop()
	inmemoryStore := services.NewInMemoryLaptopStore()
	err := inmemoryStore.Save(laptopDuplicateId)
	require.NoError(t, err)

	testCases := []struct {
		name   string
		laptop *go_grpc.Laptop
		store  services.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  services.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_without_id",
			laptop: laptopWithoutId,
			store:  services.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidId,
			store:  services.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateId,
			store:  inmemoryStore,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := go_grpc.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			service := services.NewLaptopService(tc.store)
			res, err := service.CreateLaptop(context.Background(), &req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
