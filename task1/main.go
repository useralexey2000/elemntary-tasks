package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w, h, err := readArgs()
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	arr := checks(w, h)
	print(arr)
}

func usage(n string) {
	fmt.Printf("usage: %v %v %v\n", n, "width<int>", "height<int>")
}

func checks(w, h int) [][]string {
	outer := make([][]string, 0)
	for i := 0; i < h; i++ {
		var pattern string
		if i%2 > 0 {
			pattern = " *"
		} else {
			pattern = "* "
		}
		inner := make([]string, 0)
		for j := 0; j < w; j++ {
			inner = append(inner, pattern)

		}
		outer = append(outer, inner)
	}
	return outer
}

func print(arr [][]string) {
	for _, o := range arr {
		for _, i := range o {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func readArgs() (int, int, error) {
	args := os.Args
	if len(args) < 3 {
		return 0, 0, errors.New("Arguments are < 3")
	}

	w, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, err
	}
	h, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, err
	}
	return w, h, nil
}
