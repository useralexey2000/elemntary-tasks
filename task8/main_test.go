package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name  string
		start int
		fin   int
		want  []int
	}{
		{name: "start<0", start: -1, fin: 5, want: []int{}},
		{name: "fin<0", start: 5, fin: -5, want: []int{}},
		{name: "fin<0", start: 5, fin: -5, want: []int{}},
		{name: "range0-0", start: 0, fin: 0, want: []int{0}},
		{name: "range0-3", start: 0, fin: 3, want: []int{0, 1, 1, 2}},
		{name: "range3-6", start: 3, fin: 6, want: []int{2, 3, 5, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := fib(tt.start, tt.fin)

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("%s: fib(%d, %d) = %v want %v", tt.name, tt.start, tt.fin, got, tt.want)
			}
		})
	}
}

func TestFibToStr(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want string
	}{
		{name: "range3-6-str", arr: []int{2, 3, 5, 8}, want: "2,3,5,8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := fibToStr(tt.arr)

			if tt.want != got {
				t.Errorf("%s: fibToStr(%v) = %v want %s", tt.name, tt.arr, got, tt.want)
			}
		})
	}
}

func TestReadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		min  int
		max  int
		err  error
	}{
		{name: "args!=3", args: []string{"main", "1"}, min: 0, max: 0, err: notCorrectArgsNum},
		{name: "min=NaN", args: []string{"main", "a", "5"}, min: 0, max: 0, err: strconv.ErrSyntax},
		{name: "max=NaN", args: []string{"main", "1", "b"}, min: 0, max: 0, err: strconv.ErrSyntax},
		{name: "ok", args: []string{"main", "1", "2"}, min: 1, max: 2, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			min, max, err := readArgs(tt.args)

			if min != tt.min || max != tt.max || !errors.Is(err, tt.err) {
				t.Errorf("%s: readArgs(%v) = %d, %d, %v, want %d, %d, %v", tt.name, tt.args, min, max, err, tt.min, tt.max, tt.err)
			}
		})
	}
}
