package services

import (
	"context"
	"errors"
	"fmt"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/jinzhu/copier"
	"log"
	"sync"
)

var ErrAlreadyExists = errors.New("laptop store already exists")

type LaptopStore interface {
	Save(laptop *go_grpc.Laptop) error
	Find(id string) (*go_grpc.Laptop, error)
	Search(ctx context.Context, filter *go_grpc.Filter, found func(laptop *go_grpc.Laptop) error) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*go_grpc.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data:  make(map[string]*go_grpc.Laptop),
		mutex: sync.RWMutex{},
	}
}

func (s *InMemoryLaptopStore) Save(laptop *go_grpc.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.data[laptop.Id]
	if ok {
		return ErrAlreadyExists
	}
	//deep copy
	//s.data[laptop.Id] = laptop
	//this will not work...
	copyLaptop := &go_grpc.Laptop{}
	err := copier.Copy(copyLaptop, laptop)
	if err != nil {
		return fmt.Errorf("Failed to copy Laptop with id %s: %v", laptop.Id, err)
	}
	//save the copy into the map...
	s.data[laptop.Id] = copyLaptop
	return nil
}

func (s *InMemoryLaptopStore) Find(id string) (*go_grpc.Laptop, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	laptop, ok := s.data[id]
	if !ok {
		return nil, nil
	}
	//deep copy....
	return deepCopy(laptop)
}

func (s *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *go_grpc.Filter,
	found func(laptop *go_grpc.Laptop) error,
) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, laptop := range s.data {
		//time.Sleep(1 * time.Second)
		log.Printf("checking laptop id: %v", laptop.Id)
		if errors.Is(ctx.Err(), context.DeadlineExceeded) || errors.Is(ctx.Err(), context.Canceled) {
			log.Printf("context deadline exceeded/cancelled: %v", ctx.Err())
			return errors.New("context deadline exceeded/cancelled")
		}
		if isQualified(laptop, filter) {
			//deep copy
			copyLaptop, err := deepCopy(laptop)
			if err != nil {
				return err
			}
			err = found(copyLaptop)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func isQualified(laptop *go_grpc.Laptop, filter *go_grpc.Filter) bool {
	log.Printf("checking laptop with id: %v , price: %v, cpu", laptop.GetId(), laptop.GetPriceUsd())
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if toBit(laptop.GetMemory()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func toBit(memory *go_grpc.Memory) uint64 {
	value := memory.Value

	switch memory.GetUnit() {
	case go_grpc.Memory_BIT:
		return value
	case go_grpc.Memory_BYTE:
		return value << 3 // = 2^3
	case go_grpc.Memory_KILOBYTE:
		return value << 13
	case go_grpc.Memory_MEGABYTE:
		return value << 23
	case go_grpc.Memory_GIGABYTE:
		return value << 33
	case go_grpc.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}

func deepCopy(laptop *go_grpc.Laptop) (*go_grpc.Laptop, error) {
	laptopCopy := &go_grpc.Laptop{}
	err := copier.Copy(laptopCopy, laptop)
	if err != nil {

		return nil, fmt.Errorf("Failed to copy Laptop with id %s: %v", laptop.Id, err)
	}
	return laptopCopy, nil
}
