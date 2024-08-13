package main

import (
	"testing"
)

	var lines = []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	}

// Expect these final registers
// d: 72
// e: 507
// f: 492
// g: 114
// h: 65412
// i: 65079
// x: 123
// y: 456

func Test_someAssemblyRequired(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		part  int
		want  int
	}{
		{"example h -> a", lines, 1, 65412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(&tt.input); got != tt.want {
				t.Errorf("puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}
