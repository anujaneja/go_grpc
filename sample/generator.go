package sample

import go_grpc "github.com/anujaneja/go_grpc/pb"

// NewKeyboard returns a new sample Keyboard
func NewKeyboard() *go_grpc.Keyboard {
	return &go_grpc.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

// NewCPU returns a random CPU
func NewCPU() *go_grpc.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	number_cores := randomInt(2, 8)
	number_threads := randomInt(number_cores, 16)
	min_ghz := randomFloat64(2.0, 3.5)
	max_ghz := randomFloat64(min_ghz, 5.0)

	return &go_grpc.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(number_cores),
		NumberThreads: uint32(number_threads),
		MinGhz:        min_ghz,
		MaxGhz:        max_ghz,
	}
}

func NewGPU() *go_grpc.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	min_ghz := randomFloat64(1.0, 1.5)
	max_ghz := randomFloat64(min_ghz, 2.0)
	memory := &go_grpc.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  go_grpc.Memory_GIGABYTE,
	}
	return &go_grpc.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: min_ghz,
		MaxGhz: max_ghz,
		Memory: memory,
	}
}

func NewRAM() *go_grpc.Memory {
	return &go_grpc.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  go_grpc.Memory_GIGABYTE,
	}
}

func NewSSD() *go_grpc.Storage {
	return &go_grpc.Storage{
		Driver: go_grpc.Storage_SSD,
		Memory: &go_grpc.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  go_grpc.Memory_GIGABYTE,
		},
	}
}

func NewHDD() *go_grpc.Storage {
	return &go_grpc.Storage{
		Driver: go_grpc.Storage_HDD,
		Memory: &go_grpc.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  go_grpc.Memory_GIGABYTE,
		},
	}
}

// New screen - returns the new sample screen
func NewScreen() *go_grpc.Screen {
	return &go_grpc.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
}

// New laptop
func NewLaptop() *go_grpc.Laptop {
	brand := randomLaptopBrand()
	return &go_grpc.Laptop{
		Id:       randomId(),
		Brand:    brand,
		Name:     randomLaptopName(brand),
		Cpu:      NewCPU(),
		Memory:   NewRAM(),
		Gpus:     []*go_grpc.GPU{NewGPU()},
		Storages: []*go_grpc.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &go_grpc.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: randomFloat64(2020, 2025),
	}
}
