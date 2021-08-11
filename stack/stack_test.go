package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack(2)
	if len(s) != 1 {
		t.Error("Expected length to equal 1")
	}
}
