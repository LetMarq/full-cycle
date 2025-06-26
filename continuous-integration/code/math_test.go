package main

import "testing"


func TestSum(t *testing.T) {
	total := sum(1, 2)

	if total != 3 {
		t.Errorf("Sum(1, 2) = %d; want 3", total)
	}
}