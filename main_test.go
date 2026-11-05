package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 20)

	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}
