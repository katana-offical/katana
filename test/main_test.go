package test

import "testing"

func TestMax(t *testing.T) {
	var source = 1
	var target = 2
	if source != target {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
