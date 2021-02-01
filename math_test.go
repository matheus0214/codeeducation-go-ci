package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(1, 3)

	if res != 4 {
		t.Errorf("Result is invalid! recive %d expect %d", res, 4)
	}
}
