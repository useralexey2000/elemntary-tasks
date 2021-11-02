package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var notCorrectArgsNum = errors.New("not correct number of args")

func main() {
	min, max, num, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}
	nums := naturalNums(min, max, num)
	fmt.Println(strings.Join(nums, ","))
}

// ./main min max n
func naturalNums(min, max, n int) []string {
	nums := make([]string, 0)
	// n cant be negative max cant be < min max cant be < 0
	if n <= 0 || min > max || max < 0 {
		return nums
	}

	if min < 0 {
		min = 0
	}
	// max  cant be > sqrt of n
	maxNew := int(math.Sqrt(float64(n)))
	if max > maxNew {
		max = maxNew
	}

	for i := min; i <= maxNew; i++ {
		nums = append(nums, strconv.Itoa(i))
	}
	return nums
}

func readArgs(args []string) (int, int, int, error) {

	if len(args) != 4 {
		return 0, 0, 0, notCorrectArgsNum
	}

	min, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("cant read arg min %w", err)
	}

	max, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("cant read arg max %w", err)
	}

	num, err := strconv.Atoi(args[3])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("cant read arg num %w", err)
	}

	return min, max, num, nil
}

func usage(n string) {
	fmt.Printf("usage: %v min<int> max<int> number<int>\n", n)
}
