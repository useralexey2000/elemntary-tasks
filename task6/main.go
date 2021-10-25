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

func main() {
	min, max, file, err := readArgs()
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}
	var a algo
	a, err = getAlgo(file)

	countLuckyNum(min, max, a)

}

type algo func(int, int) int

func countLuckyNum(min, max int, a algo) {
	res := a(min, max)
	fmt.Println(res)
}

func piterAlgo(min, max int) int {
	var counter int

	for i := min; i <= max; i++ {
		bt := strconv.Itoa(i)

		right := (bt[0] ^ 48) + (bt[1] ^ 48) + (bt[2] ^ 48)
		left := (bt[3] ^ 48) + (bt[4] ^ 48) + (bt[5] ^ 48)

		if right == left {
			counter++
		}

	}
	return counter
}

func moscowAlgo(min, max int) int {
	var counter int
	var even byte
	var odd byte

	for i := min; i <= max; i++ {
		bt := strconv.Itoa(i)

		for j := 0; j < len(bt); j++ {
			if bt[j]&1 == 1 {
				odd += (bt[j] ^ 48)
			} else {
				even += (bt[j] ^ 48)
			}
		}

		if even == odd {
			counter++
		}
	}
	return counter
}

func readArgs() (int, int, string, error) {
	var (
		min      int
		max      int
		filename string
	)

	args := os.Args
	if len(args) < 3 {
		return 0, 0, "", errors.New("Not enough args")
	}
	var err error
	if _, err = fmt.Scanf("%s", &filename); err != nil {
		return 0, 0, "", err
	}

	if _, err = fmt.Scanf("%d", &min); err != nil {
		return 0, 0, "", err
	}

	if _, err = fmt.Scanf("%d", &max); err != nil {
		return 0, 0, "", err
	}

	return min, max, filename, nil
}

func getAlgo(f string) (algo, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
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
	fmt.Printf("usage: %v filename<string> | filename<string> min<int> max<int>\n", n)
}
