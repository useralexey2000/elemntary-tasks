package main

import (
	"strconv"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := chess(tt.w, tt.h); got != tt.want {
				t.Errorf("%s: chess(%d, %d) = %v, want %v", tt.name, tt.w, tt.h, got, tt.want)
			}
		})
	}
}

func TestReadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		w    int
		h    int
		err  error
	}{
		{name: "wrongArgsNum", args: []string{"main"}, w: 0, h: 0, err: wrongArgsNumber},
		{name: "wrongWType", args: []string{"main", "a", "1"}, w: 0, h: 0, err: strconv.ErrSyntax},
		{name: "wrongHType", args: []string{"main", "1", "a"}, w: 0, h: 0, err: strconv.ErrSyntax},
		{name: "correct4x5", args: []string{"main", "4", "5"}, w: 4, h: 5, err: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w, h, err := readArgs(tt.args)
			if w != tt.w && h != tt.h && err != tt.err {
				t.Errorf("%s: readArgs(%v) = %d, %d, %v want: %d, %d, %v ", tt.name, tt.args, w, h, err, tt.w, tt.h, tt.err)
			}
		})
	}
}
