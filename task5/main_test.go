package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

// func readArgs(args []string) (int, error) {
func TestReadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		num  int
		err  error
	}{
		{name: "args!=2", args: []string{"main", "1", "2"}, num: 0, err: notCorrectArgsNum},
		{name: "num=NaN", args: []string{"main", "a"}, num: 0, err: strconv.ErrSyntax},
		{name: "ok", args: []string{"main", "100"}, num: 100, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			num, err := readArgs(tt.args)

			if num != tt.num || !errors.Is(err, tt.err) {
				t.Errorf("%s: readArgs(%v) = %d, %v, want %d, %v", tt.name, tt.args, num, err, tt.num, tt.err)
			}
		})
	}
}

// func hundredBlocks(i int) []int {
func TestHundredBlocks(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want []int
	}{
		{name: "i==0", i: 0, want: []int{0}},
		{name: "i>0(1001)", i: 1001, want: []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := hundredBlocks(tt.i)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: hundredBlocks(%d) = %v, want %v", tt.name, tt.i, res, tt.want)
			}
		})
	}
}

// func numBlocks(arr []int) [][]num {
func TestNumBlocks(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want [][]num
	}{
		{name: "emptyArr", arr: []int{}, want: [][]num{}},
		{name: "zero", arr: []int{0}, want: [][]num{{{0, 0, 0}}}},
		{name: "1 000 020 001", arr: []int{1, 20, 0, 1}, want: [][]num{
			{{1, 0, 3}},
			{{0, 0, 2}},
			{{0, 0, 2}},
			{{2, 2, 1}},
			{{1, 0, 0}},
		}},
		// [[{12 1 3}] [{0 0 2}] [{0 0 2}] [{2 2 1} {4 0 1}] [{5 3 0} {18 1 0}]]
		{name: "12 000 024 518", arr: []int{518, 24, 0, 12}, want: [][]num{
			{{12, 1, 3}},
			{{0, 0, 2}},
			{{0, 0, 2}},
			{{2, 2, 1}, {4, 0, 1}},
			{{5, 3, 0}, {18, 1, 0}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := numBlocks(tt.arr)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: numBlocks(%v) = %v, want %v", tt.name, tt.arr, res, tt.want)
			}
		})
	}
}

// func sprintNum(nums [][]num, m map[int]map[int]string) [][]string {
func TestSprintArrNum(t *testing.T) {

	initNumbers()

	tests := []struct {
		name string
		arr  [][]num
		want [][]string
	}{
		// {name: "emptyArr", arr: [][]num{}, want: [][]string{}},
		{name: "zero", arr: [][]num{{{0, 0, 0}}}, want: [][]string{{"ноль"}}},
		{name: "1 000 020 001", arr: [][]num{
			{{1, 0, 3}},
			{{0, 0, 2}},
			{{0, 0, 2}},
			{{2, 2, 1}},
			{{1, 0, 0}},
		},
			want: [][]string{
				{"один", "миллиард"}, {"двадцать", "тысяч"}, {"один"},
			},
		},
		{name: "1100", arr: [][]num{
			{{1, 0, 1}},
			{{1, 3, 0}},
		},
			want: [][]string{
				{"одна", "тысяча"},
				{"сто"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := sprintArrNum(tt.arr, numbers)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: numBlocks(%v) = %v, want %v", tt.name, tt.arr, res, tt.want)
			}
		})
	}
}
