package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestNaturalNums(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
		num  int
		want []string
	}{
		{name: "1num", min: 10, max: 135, num: 100, want: []string{"10"}},
		{name: "10nums", min: 8, max: 421, num: 320, want: []string{"8", "9", "10", "11", "12", "13", "14", "15", "16", "17"}},
		{name: "num=0", min: 8, max: 200, num: 0, want: []string{}},
		{name: "min<0", min: -8, max: 200, num: 120, want: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
		{name: "max<0", min: 8, max: -421, num: 320, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := naturalNums(tt.min, tt.max, tt.num)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: naturalNums(%d, %d, %d) = %v, want %v", tt.name, tt.min, tt.max, tt.num, got, tt.want)
			}
		})
	}
}

// func TestReadArgs(args []string) (int, int, int, error) {
func TestReadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		min  int
		max  int
		num  int
		err  error
	}{
		{name: "args<4", args: []string{"main", "1", "2"}, min: 0, max: 0, num: 0, err: notCorrectArgsNum},
		{name: "min=NaN", args: []string{"main", "a", "2", "3"}, min: 0, max: 0, num: 0, err: strconv.ErrSyntax},
		{name: "max=NaN", args: []string{"main", "1", "b", "3"}, min: 0, max: 0, num: 0, err: strconv.ErrSyntax},
		{name: "num=NaN", args: []string{"main", "1", "2", "c"}, min: 0, max: 0, num: 0, err: strconv.ErrSyntax},
		{name: "ok", args: []string{"main", "1", "2", "3"}, min: 1, max: 2, num: 3, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			min, max, num, err := readArgs(tt.args)

			if min != tt.min || max != tt.max || num != tt.num || !errors.Is(err, tt.err) {
				t.Errorf("%s: readArgs(%v) = %d, %d, %d, %v, want %d, %d, %d, %v", tt.name, tt.args, min, max, num, err, tt.min, tt.max, tt.num, tt.err)
			}
		})
	}
}
