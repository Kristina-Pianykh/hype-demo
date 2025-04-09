package main

import (
	"testing"
)

// TestHelloWorld is just a sanity check to make sure it doesn't panic
func TestHelloWorld(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("HelloWorld panicked: %v", r)
		}
	}()

	HelloWorld()
}

// TestAdd checks if the Add function does what it says on the tin
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{-1, -1, -2},
		{0, 0, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

// TestRandomNumber ensures the result is within range [0, 99]
func TestRandomNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := RandomNumber()
		if num < 0 || num >= 100 {
			t.Errorf("RandomNumber() = %d; out of expected range", num)
		}
	}
}
