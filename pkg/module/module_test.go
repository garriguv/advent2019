package module

import (
	"testing"
)

func TestFuelRequired(t *testing.T) {
	input := `12
1969`

	if want, got := 656, FuelRequired(input, FuelRequiredForModule); want != got {
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
		if want, got := test.want, FuelRequiredForModule(test.in); want != got {
			t.Errorf("incorrect FuelRequiredForModule: want %d got %d", want, got)
		}
	}
}

func TestFuelRequiredForModule2(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{12, 2},
		{1_969, 966},
		{100_756, 50_346},
	}
	for _, test := range tests {
		if want, got := test.want, FuelRequiredForModule2(test.in); want != got {
			t.Errorf("incorrect FuelRequiredForModule2: want %d got %d", want, got)
		}
	}
}
