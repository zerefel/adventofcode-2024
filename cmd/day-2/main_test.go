package main

import (
	"testing"
)

func TestCountSafeReports(t *testing.T) {
	t.Run("safe reports without dampener", func(t *testing.T) {
		data := []string{
			"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9",
		}

		want := 2
		got, err := CountSafeReports(data, false)

		if err != nil {
			t.Errorf("expected no error, got %s", err.Error())
		}

		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
	t.Run("safe reports with dampener", func(t *testing.T) {
		data := []string{
			"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9",
		}

		want := 4
		got, err := CountSafeReports(data, true)

		if err != nil {
			t.Errorf("expected no error, got %s", err.Error())
		}

		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("edge cases with dampener", func(t *testing.T) {
		data := []string{
			"54 53 54 55 45", // unsafe
			"54 53 55 56 57", // safe
			"1 2 3 4 5",      // Safe, increasing by 1
			"5 4 3 2 1",      // Safe, decreasing by 1
			"1 3 5 7 9",      // Safe, increasing by 2
			"9 7 5 3 1",      // Safe, decreasing by 2
			"1 4 7 10 13",    // Safe, increasing by 3
			"13 10 7 4 1",    // Safe, decreasing by 3
			"1 2 7 8 9",      // Unsafe, 2 to 7 is an increase of 5
			"9 7 6 2 1",      // Unsafe, 6 to 2 is a decrease of 4
			"1 3 2 4 5",      // Unsafe, mixed increasing and decreasing
			"8 6 4 4 1",      // Unsafe, 4 to 4 is neither increasing nor decreasing
			"1 3 6 7 9",      // Safe, increasing by 1, 2, or 3
			"1 2 3 5 8",      // Unsafe, 5 to 8 is an increase of 3
			"8 5 3 2 1",      // Safe, decreasing by 1, 2, or 3
			"1 2 3 4 6",      // Unsafe, 4 to 6 is an increase of 2
			"6 5 4 3 2",      // Safe, decreasing by 1
			"2 4 6 8 10",     // Safe, increasing by 2
			"10 8 6 4 2",     // Safe, decreasing by 2
			"1 2 3 4 7",      // Unsafe, 4 to 7 is an increase of 3
			"7 6 5 4 3",      // Safe, decreasing by 1
			"3 6 9 12 15",    // Safe, increasing by 3
		}
		want := 13
		got, err := CountSafeReports(data, true)

		if err != nil {
			t.Errorf("expected no error, got %s", err.Error())
		}

		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
