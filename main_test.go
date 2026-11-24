package main

import "testing"

func TestCanAttendAll(t *testing.T) {
	cases := []struct {
		in   [][2]int
		want bool
	}{
		{[][2]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][2]int{{7, 10}, {2, 4}}, true},
		{[][2]int{{1, 2}, {2, 3}}, true}, // back-to-back is attendable
		{[][2]int{}, true},
		{[][2]int{{5, 8}}, true},
		{[][2]int{{1, 5}, {4, 6}, {7, 8}}, false},
	}
	for _, c := range cases {
		if got := CanAttendAll(c.in); got != c.want {
			t.Errorf("CanAttendAll(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
