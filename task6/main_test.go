package main

import (
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestCheckNum(t *testing.T) {
	tests := []struct {
		name string
		num  []int
		ok   bool
	}{
		{name: "num!=6", num: []int{0, 1, 2, 2, 1, 0, 7}, ok: false},
		{name: "numIsNegative", num: []int{0, 1, 2, 2, -1, 0}, ok: false},
		{name: "numIs>10", num: []int{0, 1, 2, 2, 10, 0}, ok: false},
		{name: "numIsOk", num: []int{0, 1, 2, 2, 0, 0}, ok: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ok := checkNum(tt.num)

			if ok != tt.ok {
				t.Errorf("%s: checkNum(%v) = %t want %t", tt.name, tt.num, ok, tt.ok)
			}
		})
	}
}

// func piterAlgo(num []int) (bool, error) {
func TestPiterAlgo(t *testing.T) {
	tests := []struct {
		name  string
		args  []int
		lucky bool
		err   error
	}{
		{name: "num!=6", args: []int{0, 1, 2, 2, 1, 0, 7}, lucky: false, err: notCorrectNum},
		{name: "lucky", args: []int{0, 1, 2, 2, 1, 0}, lucky: true, err: nil},
		{name: "notlucky", args: []int{0, 1, 2, 3, 4, 5}, lucky: false, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lucky, err := piterAlgo(tt.args)

			if lucky != tt.lucky || err != tt.err {
				t.Errorf("%s: piterAlgo(%v) = %t, %v want %t, %v", tt.name, tt.args, lucky, err, tt.lucky, tt.err)
			}
		})
	}
}

func TestMoscowAlgo(t *testing.T) {
	tests := []struct {
		name  string
		args  []int
		lucky bool
		err   error
	}{
		{name: "num!=6", args: []int{0, 1, 2, 2, 1, 0, 7}, lucky: false, err: notCorrectNum},
		{name: "lucky", args: []int{0, 1, 2, 3, 4, 5}, lucky: true, err: nil},
		{name: "notlucky", args: []int{0, 1, 1, 3, 4, 5}, lucky: false, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lucky, err := moscowAlgo(tt.args)

			if lucky != tt.lucky || err != tt.err {
				t.Errorf("%s: moscowAlgo(%v) = %t, %v want %t, %v", tt.name, tt.args, lucky, err, tt.lucky, tt.err)
			}
		})
	}
}

// func countLuckyNum(nums [][]int, al algo) (int, error) {
func TestCountLuckyNum(t *testing.T) {
	tests := []struct {
		name  string
		nums  [][]int
		algo  algo
		count int
		err   error
	}{
		{name: "piterlucky=1", nums: [][]int{{1, 1, 1, 0, 0, 3}, {0, 0, 0, 1, 0, 0}}, algo: piterAlgo, count: 1, err: nil},
		{name: "moscowlucky=2", nums: [][]int{{1, 1, 1, 0, 0, 0}, {2, 2, 2, 1, 1, 1}}, algo: moscowAlgo, count: 2, err: nil},
		{name: "err=negativenum", nums: [][]int{{1, 1, -1, 0, 0, 0}, {2, 2, 2, 1, 1, 1}}, algo: moscowAlgo, count: 0, err: notCorrectNum},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			count, err := countLuckyNum(tt.nums, tt.algo)

			if !errors.Is(err, tt.err) || count != tt.count {
				t.Errorf("%s: CountLuckyNum(%v, %v) = %d, %v want %d, %v", tt.name, tt.nums, tt.algo, count, err, tt.count, tt.err)
			}
		})
	}
}

func TestReadArgs(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		fname string
		nums  [][]int
		err   error
	}{
		{name: "args!=3", args: []string{"main", "fname"}, fname: "", nums: nil, err: argLenghtErr},
		{name: "argsNaN", args: []string{"main", "fname", "01234a"}, fname: "", nums: nil, err: strconv.ErrSyntax},
		{name: "argsok", args: []string{"main", "fname", "012345"}, fname: "fname", nums: [][]int{{0, 1, 2, 3, 4, 5}}, err: nil},
		{name: "argsMultNums", args: []string{"main", "fname", "012345", "543210"}, fname: "fname", nums: [][]int{{0, 1, 2, 3, 4, 5}, {5, 4, 3, 2, 1, 0}}, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fname, nums, err := readArgs(tt.args)

			if fname != tt.fname || !reflect.DeepEqual(nums, tt.nums) || !errors.Is(err, tt.err) {
				t.Errorf("%s: readArgs(%v) = %s, %v, %v, want %s, %v, %v", tt.name, tt.args, fname, nums, err, tt.fname, tt.nums, tt.err)
			}
		})
	}
}

// func getAlgo(r io.Reader) (algo, bool) {
func TestGetAlgo(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		al     algo
		found  bool
	}{
		{name: "piterAlgo", reader: strings.NewReader(piter), al: piterAlgo, found: true},
		{name: "moscowAlgo", reader: strings.NewReader(moscow), al: moscowAlgo, found: true},
		{name: "notFound", reader: strings.NewReader("unknownAlgo"), al: nil, found: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			al, found := getAlgo(tt.reader)
			alEqual := reflect.ValueOf(al).Pointer() == reflect.ValueOf(tt.al).Pointer()
			if found != tt.found || !alEqual {
				t.Errorf("%s: getAlgo(reader) = %v, %t want: %v, %t", tt.name, al, found, tt.al, tt.found)

			}
		})
	}
}
