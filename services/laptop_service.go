package services

import (
	context "context"
	"errors"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type LaptopService struct {
	go_grpc.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func NewLaptopService(store LaptopStore) *LaptopService {
	return &LaptopService{
		Store: store,
	}
}
func (service *LaptopService) CreateLaptop(
	ctx context.Context,
	request *go_grpc.CreateLaptopRequest,
) (*go_grpc.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("recieved a create-laptop request for laptop id %v", laptop.Id)

	if len(laptop.Id) > 0 {
		//check if its valid uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid laptop id %v", laptop.Id)
		}
	} else {
		//generate a new id
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not create laptop id %v", laptop.Id)
		}
		laptop.Id = id.String()
	}

	/**
	Code for putting a sleep to check for timeout...
	*/
	time.Sleep(6 * time.Second)
	if errors.Is(ctx.Err(), context.Canceled) {
		log.Printf("request is cancelled for laptop id %v", laptop.Id)
		return nil, status.Errorf(codes.Canceled, "context canceled")
	}
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Printf("timed out waiting for laptop id %v", laptop.Id)
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	//Save it to the in-memory store...
	err := service.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "could not save laptop %v, %v", laptop.Id, err)
	}
	log.Printf("Laptop saved with id %v", laptop.Id)
	res := &go_grpc.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
