package main

import "testing"

func TestCalculateDistance(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}
	expected := 11
	got := CalculateTotalDistance(a, b)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}

}

func TestCalculateSimilarity(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}
	expected := 31
	got := CalculateSimilarty(a, b)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
