package module

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func FuelRequired(input string, fn func(int) int) int {
	fuel := 0
	for _, in := range strings.Split(input, "\n") {
		if in == "" {
			continue
		}

		module, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal("invalid input:", err)
		}
		fuel += fn(module)
	}
	return fuel
}

func FuelRequiredForModule(mass int) int {
	return fuel(mass)
}

func FuelRequiredForModule2(mass int) int {
	var (
		f   = fuel(mass)
		acc = f
	)
	for {
		if f = fuel(f); f <= 0 {
			break
		}
		acc += f
	}
	return acc
}

func fuel(mass int) int {
	return int(math.Floor(float64(mass)/3) - 2)
}
