package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

// func readArgs(args []string) (int, error)
func TestReadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		num  int
		err  error
	}{
		{name: "args!=2", args: []string{"main", "1", "2"}, num: 0, err: errArgsNum},
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

// func splitThousand(i int) []int
func TestSplitThousand(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want []int
	}{
		{name: "i==0", i: 0, want: []int{0}},
		{name: "i==(12 000 123 002)", i: 12000123002, want: []int{12, 0, 123, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := splitThousand(tt.i)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: splitHundred(%d) = %v, want %v", tt.name, tt.i, res, tt.want)
			}
		})
	}
}

// func splitHundred(i int) []int
func TestSplitHundred(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want []int
	}{
		{name: "i==0", i: 0, want: []int{0, 0, 0}},
		{name: "i>0(121)", i: 121, want: []int{1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := splitHundred(tt.i)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: splitHundred(%d) = %v want %v", tt.name, tt.i, res, tt.want)
			}
		})
	}
}

// func constructNum(i int) (*Num, error)
func TestConstructNum(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want *Num
	}{
		{name: "i<0", i: -1, want: &Num{positive: false, val: [][]int{{0, 0, 1}}}},
		{name: "i=210 102", i: 210102,
			want: &Num{
				positive: true,
				val: [][]int{
					{2, 0, 10},
					{1, 0, 2},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := constructNum(tt.i)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: splitNumber(%d) = %v, want %v", tt.name, tt.i, res, tt.want)
			}
		})
	}
}

// func numToArrText(num *Num, m map[int]map[int]string) [][]string {
func TestNumToArrText(t *testing.T) {

	numbers := initNumberMapper()

	tests := []struct {
		name string
		num  *Num
		want [][]string
	}{
		{name: "zero", num: &Num{positive: true, val: [][]int{{0, 0, 0}}}, want: [][]string{{"ноль"}}},

		{name: "minus", num: &Num{positive: false, val: [][]int{{1, 2, 2}}},
			want: [][]string{{"минус"}, {"сто", "двадцать", "два"}}},

		{name: "1000", num: &Num{positive: true, val: [][]int{{0, 0, 1}, {0, 0, 0}}},
			want: [][]string{
				{"одна", "тысяча"}}},

		{name: "1 000 020 001", num: &Num{
			positive: true,
			val: [][]int{
				{0, 0, 1},
				{0, 0, 0},
				{0, 2, 0},
				{0, 0, 1}}},
			want: [][]string{
				{"один", "миллиард"}, {"двадцать", "тысяч"}, {"один"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := numToArrText(tt.num, numbers)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: numToArrText(%v) = %v, want %v", tt.name, tt.num, res, tt.want)
			}
		})
	}
}
