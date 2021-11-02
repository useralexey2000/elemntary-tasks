package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	moscow = "Moskow"
	piter  = "Piter"
)

var notCorrectNum = errors.New("cot correct number")
var argLenghtErr = errors.New("wrong length of args, must  be > 2")

func main() {
	// readArgs([]string{"main", "fname", "012345", "543210"})

	file, nums, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(file)
	if err != nil {
		panic("cant open file")
	}
	al, found := getAlgo(f)

	if !found {
		panic("cant find algo")
	}

	res, err := countLuckyNum(nums, al)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	fmt.Println("Number of lucky numbers :", res)

}

type algo func([]int) (bool, error)

func countLuckyNum(nums [][]int, al algo) (int, error) {
	var counter int

	for _, v := range nums {
		b, err := al(v)
		if err != nil {
			return 0, err
		}
		if b {
			counter++
		}
	}

	return counter, nil
}

// check if len == 6
// check if negative
// check if n > 9
func checkNum(num []int) bool {
	if len(num) != 6 {
		return false
	}

	for _, v := range num {
		if v > 9 || v < 0 {
			return false
		}
	}
	return true
}

func piterAlgo(num []int) (bool, error) {

	if !checkNum(num) {
		return false, notCorrectNum
	}

	right := num[0] + num[1] + num[2]
	left := num[3] + num[4] + num[5]

	return left == right, nil
}

func moscowAlgo(num []int) (bool, error) {

	if !checkNum(num) {
		return false, notCorrectNum
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

// returns filename for algo, nested array of ints each lengh=6(length of ticket), err
func readArgs(args []string) (string, [][]int, error) {
	// at least one number should be specified
	if len(args) < 3 {
		return "", nil, argLenghtErr
	}

	fname := args[1]
	strNums := args[2:]

	nums := make([][]int, 0)

	for _, s := range strNums {

		num := make([]int, 0)
		arr := strings.Split(s, "")
		for _, v := range arr {

			n, err := strconv.Atoi(v)
			if err != nil {
				return "", nil, fmt.Errorf("number cant be processed %w", err)
			}
			num = append(num, n)
		}
		nums = append(nums, num)
	}

	return fname, nums, nil

}

func getAlgo(r io.Reader) (algo, bool) {

	scanner := bufio.NewScanner(r)

	scanner.Scan()

	str := scanner.Text()

	if strings.Contains(str, moscow) {
		return moscowAlgo, true
	}

	if strings.Contains(str, piter) {
		return piterAlgo, true
	}

	return nil, false
}

func usage(n string) {
	fmt.Printf("usage: %v filename<string> nums<int>...\n", n)
}
