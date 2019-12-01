package module

import (
	"testing"
)

func TestFuelRequired(t *testing.T) {
	input := `12
1969`

	if want, got := 656, FuelRequired(input); want != got {
		t.Errorf("incorrect FuelRequired: want %d got %d", want, got)
	}
}

func TestFuelRequiredForModule(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1_969, 654},
		{100_756, 33_583},
	}
	for _, test := range tests {
		if want, got := test.want, fuelRequiredForModule(test.in); want != got {
			t.Errorf("incorrect fuelRequired: want %d got %d", want, got)
		}
	}
}
