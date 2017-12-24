package string

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		forward, backward string
	}{
		{"Backward", "drawkcaB"},
		{"Hello,", ",olleH"},
	}
	for _, c := range tests {
		returnValue := Reverse(c.forward)
		if returnValue != c.backward {
			t.Errorf("Reverse(%q) == %q, want %q", c.forward, returnValue, c.backward)
		}
	}
}
