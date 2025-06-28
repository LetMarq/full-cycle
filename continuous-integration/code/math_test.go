package main

import "testing"

func TestSum(t *testing.T) {
	total := sumFunction(1, 2)

	if total != 3 {
		t.Errorf("sumFunction(1, 2) = %d; want 3", total)
	}
}
