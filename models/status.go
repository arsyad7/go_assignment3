package models

import (
	"math/rand"
	"time"
)

type Status struct {
	Wind  Wind
	Water Water
}

func GetStatus() Status {
	var result Status

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	result.Water.Status = random.Intn(100)
	result.Wind.Status = random.Intn(100)
	result.Water.Message = GetWarningWater(result.Water.Status)
	result.Wind.Message = GetWarningWind(result.Wind.Status)

	return result
}

func GetWarningWater(num int) string {
	var result string

	switch {
	case num < 5:
		result = "aman"
	case num <= 8:
		result = "siaga"
	case num > 8:
		result = "bahaya"
	}

	return result
}

func GetWarningWind(num int) string {
	var result string

	switch {
	case num < 6:
		result = "aman"
	case num <= 15:
		result = "siaga"
	case num > 15:
		result = "bahaya"
	}

	return result
}
