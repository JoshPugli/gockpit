package numbers

import (
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		hasError bool
	}{
		{"empty", []int{}, 0, true},
		{"single value", []int{5}, 5, false},
		{"positive values", []int{1, 3, 5, 2, 4}, 5, false},
		{"negative values", []int{-1, -3, -5, -2, -4}, -1, false},
		{"mixed values", []int{-10, 0, 10, 5, -5}, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Max(tt.input...)
			if (err != nil) != tt.hasError {
				t.Errorf("Max() error = %v, wantErr %v", err, tt.hasError)
				return
			}
			if err == nil && result != tt.expected {
				t.Errorf("Max() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		hasError bool
	}{
		{"empty", []int{}, 0, true},
		{"single value", []int{5}, 5, false},
		{"positive values", []int{1, 3, 5, 2, 4}, 1, false},
		{"negative values", []int{-1, -3, -5, -2, -4}, -5, false},
		{"mixed values", []int{-10, 0, 10, 5, -5}, -10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Min(tt.input...)
			if (err != nil) != tt.hasError {
				t.Errorf("Min() error = %v, wantErr %v", err, tt.hasError)
				return
			}
			if err == nil && result != tt.expected {
				t.Errorf("Min() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		hasError bool
	}{
		{"empty", []int{}, 0, true},
		{"single value", []int{5}, 5, false},
		{"positive values", []int{1, 2, 3, 4, 5}, 15, false},
		{"negative values", []int{-1, -2, -3, -4, -5}, -15, false},
		{"mixed values", []int{-5, -2, 0, 3, 4}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Sum(tt.input...)
			if (err != nil) != tt.hasError {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.hasError)
				return
			}
			if err == nil && result != tt.expected {
				t.Errorf("Sum() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected float64
		hasError bool
	}{
		{"empty", []int{}, 0, true},
		{"single value", []int{5}, 5.0, false},
		{"positive values", []int{1, 2, 3, 4, 5}, 3.0, false},
		{"negative values", []int{-1, -2, -3, -4, -5}, -3.0, false},
		{"mixed values", []int{-5, 5}, 0.0, false},
		{"decimal result", []int{1, 2}, 1.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Average(tt.input...)
			if (err != nil) != tt.hasError {
				t.Errorf("Average() error = %v, wantErr %v", err, tt.hasError)
				return
			}
			if err == nil && result != tt.expected {
				t.Errorf("Average() = %v, want %v", result, tt.expected)
			}
		})
	}
}