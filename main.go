package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/garriguv/advent2019/pkg/module"
)

var (
	day  = flag.Int("day", 0, "day")
	days = make(map[int]Day)
)

type Part func(string) interface{}
type Day []Part

func init() {
	days[1] = Day([]Part{
		func(input string) interface{} { return module.FuelRequired(input) },
	})
}

func main() {
	flag.Parse()
	if *day <= 0 {
		log.Fatalf("invalid day parameter: %d", *day)
	}

	bytes, err := ioutil.ReadFile(fmt.Sprintf("inputs/day%d.txt", *day))
	if err != nil {
		log.Fatal("error loading input:", err)
	}

	solution, ok := days[*day]
	if !ok {
		log.Fatalf("no solution for day %d", *day)
	}

	fmt.Printf("Day %d: https://adventofcode.com/2019/day/%d\n", *day, *day)

	input := string(bytes)
	for part, fn := range solution {
		fmt.Printf("Part %d: %v\n", part+1, fn(input))
	}
}
