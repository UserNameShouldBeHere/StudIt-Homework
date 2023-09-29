package main

import "testing"

func TestSum(t *testing.T) {
	data := [][]int{
		{0, 0},
		{1, -1},
		{3, 4},
	}
	for _, row := range data {
		if row[0]+row[1] != Sum(row[0], row[1]) {
			t.Errorf("Expected: %d; Get: %d;", row[0]+row[1], Sum(row[0], row[1]))
		}
	}
}
