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

// func splitHundred(rank, i int) *NumBlock {
func TestSplitHundred(t *testing.T) {
	tests := []struct {
		name string
		rank int
		i    int
		want *NumBlock
	}{
		{name: "i==0", rank: 0, i: 0, want: &NumBlock{rank: 0, val: [3]int{0, 0, 0}}},
		{name: "i>0(121)", rank: 0, i: 121, want: &NumBlock{rank: 0, val: [3]int{1, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := splitHundred(tt.rank, tt.i)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: splitHundred(%d, %d) = %v want %v", tt.name, tt.rank, tt.i, res, tt.want)
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
		{name: "i=-1", i: -1,
			want: &Num{positive: false, block: []*NumBlock{{rank: 0, val: [3]int{0, 0, 1}}}},
		},
		{name: "i=124001", i: 124001,
			want: &Num{positive: true,
				block: []*NumBlock{
					{rank: 1, val: [3]int{1, 2, 4}},
					{rank: 0, val: [3]int{0, 0, 1}},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := constructNum(tt.i)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("%s: constructNumber(%d) = %v, want %v", tt.name, tt.i, res, tt.want)
			}
		})
	}
}

// func NumToText(i int, mapper *NumMapper) string
func TestNumToText(t *testing.T) {
	mapper := initNumberMapper()
	tests := []struct {
		name string
		i    int
		want string
	}{
		{name: "i=0", i: 0, want: "ноль"},
		{name: "i=1 000", i: 1000, want: "одна тысяча"},
		{name: "i=1 001", i: 1001, want: "одна тысяча один"},
		{name: "i=1 000 102 418", i: 1000102418,
			want: "один миллиард сто две тысячи четыреста восемнадцать"},
		{name: "i=-1 000 002 000", i: -1000005002,
			want: "минус один миллиард пять тысяч два"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := NumToText(tt.i, mapper)

			if res != tt.want {
				t.Errorf("%s: NumToText(%d, mapper) = %v want %v", tt.name, tt.i, res, tt.want)
			}
		})
	}
}
