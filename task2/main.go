package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var a, b, c, d float64
		var err error

		fmt.Println("please enter envelope #1 side a<float>")
		if a, err = readArgs(scanner); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("please enter envelope #1 side b<float>")
		if b, err = readArgs(scanner); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("please enter envelope #2 side a<float>")
		if c, err = readArgs(scanner); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("please enter envelope #2 side b<float>")
		if d, err = readArgs(scanner); err != nil {
			fmt.Println(err)
			continue
		}

		e1 := Envelope{A: a, B: b}
		e2 := Envelope{A: c, B: d}
		e1.Fit(e2)

		fmt.Println("Do you want to start again?<yes|y>")
		var s string
		if _, err := fmt.Scanln(&s); err != nil {
			fmt.Println("can`t read argument", err)
			os.Exit(1)
		}
		s = strings.ToLower(strings.Trim(s, " "))
		if s == "yes" || s == "y" {
			continue
		}
		os.Exit(0)
	}
}

type Envelope struct {
	A float64
	B float64
}

func (e Envelope) Fit(e1 Envelope) bool {
	if e.A > e1.A && e.B > e1.B || e.A > e1.B && e.B > e1.A {
		return true
	}
	return false
}

func print(e, e1 Envelope, b bool) {
	if b {
		fmt.Printf("envelope #1: (%v, %v) can accomodate envelope #2: (%v, %v)\n", e.A, e.B, e1.A, e1.B)
		return
	}
	fmt.Printf("envelope #1: (%v, %v) can't accomodate envelope #2: (%v, %v)\n", e.A, e.B, e1.A, e1.B)
}

func readArgs(scanner *bufio.Scanner) (float64, error) {

	scanner.Scan()
	str := scanner.Text()

	side, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, fmt.Errorf("can`t read argument %v", err)
	}
	return side, nil
}
