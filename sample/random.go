package sample

import (
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/google/uuid"
	"math/rand"
)

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomKeyboardLayout() go_grpc.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return go_grpc.Keyboard_QWERTY
	case 2:
		return go_grpc.Keyboard_QWERTZ
	default:
		return go_grpc.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS")
	default:
		return randomStringFromSet("Thinkpad Xi", "Thinkpad P1")
	}
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"2.16.1 Celeron (Sandy Bridge/Ivy Bridge microarchitecture)",
			"2.16.2 Pentium (Sandy Bridge/Ivy Bridge microarchitecture)",
			"2.16.3 Core i3 (2nd and 3rd generation)",
			"2.16.4 Core i5 (2nd and 3rd generation)",
			"2.16.5 Core i7 (2nd and 3rd generation))",
		)
	}
	return randomStringFromSet(
		"Ryzen 5 7600X",
		"Ryzen 7 7800X3D",
		"Ryzen 7 9700X",
		"Ryzen 5 8600G",
		"Ryzen 7 8700G.",
	)

}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GeForce RTX 3080",
			"GeForce RTX 4090",
			"GeForce RTX 4070 Ti",
			"NVIDIA RTX A4000",
		)
	}
	return randomStringFromSet(
		"Radeon RX 7900 XTX",
		"Radeon RX 7900 GRE", "Radeon RX 9070",
		"Radeon RX 9070 XT",
	)
}

func randomStringFromSet(list ...string) string {
	length := len(list)
	if length == 0 {
		return ""
	}
	return list[rand.Intn(length)]
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() go_grpc.Screen_Panel {
	if rand.Intn(2) == 1 {
		return go_grpc.Screen_IPS
	}
	return go_grpc.Screen_OLED
}

func randomScreenResolution() *go_grpc.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	return &go_grpc.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
}

func randomId() string {
	return uuid.New().String()
}
