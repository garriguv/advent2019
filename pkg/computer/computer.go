package computer

import (
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

func Compute(program string) int {
	input := parseProgram(program)

	input[1] = 12
	input[2] = 2

	runIntCode(&input)

	return input[0]
}

func FindNounVerb(program string, want int) int {
	input := parseProgram(program)
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			cpy := make([]int, len(input))
			copy(cpy, input)
			cpy[1] = noun
			cpy[2] = verb
			runIntCode(&cpy)
			if cpy[0] == want {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func runIntCode(input *[]int) {
	tmp := *input

	for i := 0; i < len(tmp); {
		opcode := tmp[i]
		switch opcode {
		case 1:
			in1 := tmp[i+1]
			in2 := tmp[i+2]
			out := tmp[i+3]
			v := tmp[in1] + tmp[in2]
			tmp[out] = v
		case 2:
			in1 := tmp[i+1]
			in2 := tmp[i+2]
			out := tmp[i+3]
			v := tmp[in1] * tmp[in2]
			tmp[out] = v
		case 99:
			goto out
		default:
			log.Fatal("invalid opcode:", opcode)
		}
		i += 4
	}

out:
	*input = tmp
}

func parseProgram(input string) []int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<',' | 1<<' ' | 1<<'\n'

	var output = make([]int, 0)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		val, err := strconv.Atoi(s.TokenText())
		if err != nil {
			log.Fatal("invalid token:", err)
		}
		output = append(output, val)
	}

	return output
}
