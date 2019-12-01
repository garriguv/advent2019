package module

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func FuelRequired(input string) int {
	fuel := 0
	for _, in := range strings.Split(input, "\n") {
		if in == "" {
			continue
		}

		module, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal("invalid input:", err)
		}
		fuel += fuelRequiredForModule(module)
	}
	return fuel
}

func fuelRequiredForModule(mass int) int {
	return int(math.Floor(float64(mass)/3) - 2)
}
