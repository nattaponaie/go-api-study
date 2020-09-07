package main

import "testing"

func TestPlus(t *testing.T) {
	res := plus(1, 2)
	if res != 3 {
		t.Errorf("Expect 3 but got %d", res)
	}
}
