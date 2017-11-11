package maths

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var slice []float64
	slice = append(slice, 3.2)
	slice = append(slice, 4.7)
	slice = append(slice, 56.7)
	slice = append(slice, 144.14)
	result := Average(slice)
	if result != 52.185 {
		t.Error("Expected 52.185, got ", result)
	}
}

func TestCalculateSuggestedPrice(t *testing.T) {
	var slice []float64
	slice = append(slice, 4.7)
	slice = append(slice, 56.7)
	slice = append(slice, 3.2)
	slice = append(slice, 144.14)
	max, suggested, min, _ := CalculateSuggestedPrice(slice)
	if min != 3.2 {
		t.Error("Expected 3.2, got ", min)
	}

	if suggested != 52 {
		t.Error("Expected 52.185, got ", suggested)
	}

	if max != 144.14 {
		t.Error("Expected 144.14, got ", max)
	}
}

func TestCalculateSuggestedPriceFail(t *testing.T) {
	var slice []float64
	_, _, _, err := CalculateSuggestedPrice(slice)
	if err == nil {
		t.Error("Expected empty slice, got ", err)
	}
}

func TestRound(t *testing.T) {
	suggested := 144.447765
	result := Round(suggested, 0.5)
	if result != 144.5 {
		t.Error("Expected 52.185, got ", result)
	}

	suggested = 144.947765
	result = Round(suggested, 0.5)
	if result != 145 {
		t.Error("Expected 52.185, got ", result)
	}

	suggested = 144.047765
	result = Round(suggested, 0.5)
	if result != 144 {
		t.Error("Expected 52.185, got ", result)
	}
}

func BenchmarkRound(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Round(float64(n), 0.5)
	}
}

func BenchmarkCalculateSuggestedPrice(b *testing.B) {
	var slice []float64
	slice = append(slice, 4.7)
	slice = append(slice, 56.7)
	slice = append(slice, 3.2)
	slice = append(slice, 144.14)
	for n := 0; n < b.N; n++ {
		CalculateSuggestedPrice(slice)
	}
}

func BenchmarkAverage(b *testing.B) {
	var slice []float64
	slice = append(slice, 4.7)
	slice = append(slice, 56.7)
	slice = append(slice, 3.2)
	slice = append(slice, 144.14)
	for n := 0; n < b.N; n++ {
		Average(slice)
	}
}
