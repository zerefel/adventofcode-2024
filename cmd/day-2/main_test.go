package main

import (
	"testing"
)

func TestCountSafeReports(t *testing.T) {
	// data := []string{
	// 	"7 6 4 2 1",
	// 	"1 2 7 8 9",
	// 	"9 7 6 2 1",
	// 	"1 3 2 4 5",
	// 	"8 6 4 4 1",
	// 	"1 3 6 7 9",
	// }

	data := []string{
		"44 47 50 51 53 54 53",
		"70 73 75 77 80 81 84 84",
		"1 3 4 7 10 13 16 20",
		"47 49 52 53 55 57 60 65",
	}
	want := 2
	got, err := CountSafeReports(data)

	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
