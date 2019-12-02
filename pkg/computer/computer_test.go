package computer

import (
	"reflect"
	"testing"
)

func TestCompute(t *testing.T) {
	input := `1,1,1,0,99,5 ,6,0,99,1,1,1,1
`

	got := Compute(input)

	if want, got := 3, got; want != got {
		t.Errorf("incorrect Compute result: want %v got %v", want, got)
	}
}

func TestRunIntCode(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}
	for _, tt := range tests {
		runIntCode(&tt.input)
		if want, got := tt.want, tt.input; !reflect.DeepEqual(want, got) {
			t.Errorf("incorrect Run result: want %v got %v", want, got)
		}
	}
}
