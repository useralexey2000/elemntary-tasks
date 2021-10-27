package main

import (
	"testing"
)

func TestChess(t *testing.T) {
	tests := []struct {
		name string
		w    int
		h    int
		want string
	}{
		{name: "zeroheight", w: 8, h: 0, want: ""},
		{name: "zerowidth", w: 0, h: 0, want: ""},
		{name: "zeroheightwidth", w: 0, h: 0, want: ""},
		{name: "8x5", w: 8, h: 5, want: "*_*_*_*_\n_*_*_*_*\n*_*_*_*_\n_*_*_*_*\n*_*_*_*_"},
		{name: "mustfail", w: 8, h: 6, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := chess(tt.w, tt.h); got != tt.want {
				t.Errorf("%s: chess(%d, %d) = %v, want %v", tt.name, tt.w, tt.h, got, tt.want)
			}
		})
	}
}
