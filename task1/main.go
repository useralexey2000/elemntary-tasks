package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	if len(args) < 3 {
		usage(args[0])
		os.Exit(1)
	}

	w, err := strconv.Atoi(args[1])
	if err != nil {
		usage(args[0])
		os.Exit(1)
	}
	h, err := strconv.Atoi(args[2])
	if err != nil {
		usage(args[0])
		os.Exit(1)
	}
	checks(w, h)
}

func usage(n string) {
	fmt.Printf("usage: %v %v %v\n", n, "width<int>", "height<int>")
}

func checks(w, h int) {
	for i := 0; i < h; i++ {
		var pattern string
		if i%2 > 0 {
			pattern = " *"
		} else {
			pattern = "* "
		}
		for j := 0; j < w; j++ {
			fmt.Print(pattern)
		}
		fmt.Println()
	}
}
