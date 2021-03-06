package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	black string = "*"
	white string = "_"
)

var wrongArgsNumber = errors.New("Arguments are != 3")

func main() {

	w, h, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	str := chess(w, h)
	fmt.Println(str)

}

func usage(n string) {
	fmt.Printf("usage: width<int> height<int> %v\n", n)
}

func chess(w, h int) string {
	if h == 0 || w == 0 {
		return ""
	}

	var sb strings.Builder

	bl, wt := black, white
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j%2 > 0 {
				//  always returns nil error
				sb.WriteString(wt)
				continue
			}

			sb.WriteString(bl)
		}
		//  dont append last blank line
		if i+1 < h {
			sb.WriteString("\n")
			bl, wt = wt, bl
		}
	}

	return sb.String()
}

func readArgs(args []string) (int, int, error) {

	if len(args) != 3 {
		return 0, 0, wrongArgsNumber
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
