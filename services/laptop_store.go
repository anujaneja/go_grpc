package services

import (
	"errors"
	"fmt"
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/jinzhu/copier"
	"sync"
)

var ErrAlreadyExists = errors.New("laptop store already exists")

type LaptopStore interface {
	Save(laptop *go_grpc.Laptop) error
	Find(id string) (*go_grpc.Laptop, error)
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
	laptopCopy := &go_grpc.Laptop{}
	err := copier.Copy(laptopCopy, laptop)
	if err != nil {
		return nil, fmt.Errorf("Failed to copy Laptop with id %s: %v", laptop.Id, err)
	}
	return laptopCopy, nil
}
