package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	moscow = "Moskow"
	piter  = "Piter"
)

var numLengthErr = errors.New("wrong length of num, must  be 6")
var argLenghtErr = errors.New("wrong length of args, must  be > 2")

func main() {
	file, nums, err := readArgs(os.Args)
	// file, nums, err := readArgs([]string{"main", "fname", "101110"})
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	a, err := getAlgo(file)

	res, err := countLuckyNum(nums, a)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	fmt.Println("Number of lucky numbers :", res)

}

type algo func([]int) (bool, error)

func countLuckyNum(nums [][]int, a algo) (int, error) {
	var counter int

	for _, v := range nums {
		b, err := a(v)
		if err != nil {
			return 0, err
		}
		if b {
			counter++
		}
	}

	return counter, nil
}

func piterAlgo(num []int) (bool, error) {
	if len(num) != 6 {
		return false, numLengthErr
	}

	right := num[0] + num[1] + num[2]
	left := num[3] + num[4] + num[5]

	return left == right, nil
}

func moscowAlgo(num []int) (bool, error) {
	if len(num) != 6 {
		return false, numLengthErr
	}

	even, odd := 0, 0

	for _, v := range num {
		if v%2 == 0 {
			even++
			continue
		}
		odd++
	}

	return even == odd, nil
}

func readArgs(args []string) (string, [][]int, error) {
	// at least one number should be specified
	if len(args) < 3 {
		return "", nil, argLenghtErr
	}

	fname := args[1]
	strNums := args[2:]

	nums := make([][]int, 0)
	num := make([]int, 0)

	for _, s := range strNums {

		arr := strings.Split(s, "")
		for _, v := range arr {

			n, err := strconv.Atoi(v)
			if err != nil {
				return "", nil, fmt.Errorf("number cant be processed %v", err)
			}
			num = append(num, n)
		}
		nums = append(nums, num)
	}

	return fname, nums, nil

}

func getAlgo(f string) (algo, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("cant open file: %s %v", f, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, fmt.Errorf("cant read from file: %s", f)
	}

	str := scanner.Text()

	if strings.Contains(str, moscow) {
		return moscowAlgo, nil
	}

	if strings.Contains(str, piter) {
		return piterAlgo, nil
	}
	return nil, fmt.Errorf("cant find algos from file: %s", f)

}

func usage(n string) {
	fmt.Printf("usage: %v filename<string> nums<int>...\n", n)
}
