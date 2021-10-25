package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	min, max, num, err := readArgs()
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}
	naturalNums(min, max, num)
}

// ./main min max n
func naturalNums(min, max, n int) {
	// max  cant be > sqrt of n
	maxNew := int(math.Sqrt(float64(n)))

	for i := min; i <= maxNew; i++ {
		fmt.Println(i)
	}
}

func readArgs() (int, int, int, error) {
	var (
		min int
		max int
		num int
	)

	args := os.Args
	if len(args) < 4 {
		return 0, 0, 0, errors.New("Not enough args")
	}
	var err error

	if _, err = fmt.Scanf("%d", &min); err != nil {
		return 0, 0, 0, err
	}

	if _, err = fmt.Scanf("%d", &max); err != nil {
		return 0, 0, 0, err
	}

	if _, err = fmt.Scanf("%s", &num); err != nil {
		return 0, 0, 0, err
	}

	return min, max, num, nil
}

func usage(n string) {
	fmt.Printf("usage: %v min<int> max<int> number<int>\n", n)
}
