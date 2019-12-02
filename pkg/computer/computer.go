package computer

import (
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

func Compute(program string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(program))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<',' | 1<<' ' | 1<<'\n'

	var input = make([]int, 0)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		val, err := strconv.Atoi(s.TokenText())
		if err != nil {
			log.Fatal("invalid token:", err)
		}
		input = append(input, val)
	}

	input[1] = 12
	input[2] = 2

	runIntCode(&input)

	return input[0]
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
