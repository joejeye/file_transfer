package myutils

import "testing"

func TestRandName(t *testing.T) {
	randName := RandName()
	t.Log(randName)
}
