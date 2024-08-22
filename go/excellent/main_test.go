package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	reult := EvenOrOdd(2)
	if reult != "Even" {
		t.Errorf("EvenOrOdd(2) = %s; want Even", reult)
	}
}