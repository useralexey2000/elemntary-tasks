package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var notCorrectArgsNum = errors.New("not correct number of args")

func main() {
	min, max, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
	}
	arr := fib(min, max)
	fmt.Println(fibToStr(arr))
}

// 0 1 1 2 3 5 8 13
func fib(start, fin int) []int {
	if start < 0 || fin < 0 {
		return []int{}
	}

	a, b := 0, 1
	// find previous number for start point
	for i := 0; i < start; i++ {
		a, b = b, a+b
	}
	nums := make([]int, 0)
	// append start
	nums = append(nums, a)
	// print remaining numbers
	for i := start; i < fin; i++ {
		a, b = b, a+b
		nums = append(nums, a)
	}
	return nums
}

func fibToStr(arr []int) string {
	str := make([]string, 0)

	for _, v := range arr {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, ",")
}

func readArgs(args []string) (int, int, error) {

	if len(args) != 3 {
		return 0, 0, notCorrectArgsNum
	}
	min, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, fmt.Errorf("cant read arg min %w", err)
	}

	max, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, fmt.Errorf("cant read arg max %w", err)
	}

	return min, max, nil
}

func usage(n string) {
	fmt.Printf("usage: %v min<int> max<int>\n", n)
}
