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

		fmt.Println("please enter envelope #1 side a<int>")
		scanner.Scan()
		str := scanner.Text()
		var err error
		if a, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println("can`t read argument", err)
			continue
		}

		fmt.Println("please enter envelope #1 side b<int>")
		scanner.Scan()
		str = scanner.Text()
		if b, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println("can`t read argument", err)
			continue
		}

		fmt.Println("please enter envelope #2 side a<int>")
		scanner.Scan()
		str = scanner.Text()
		if c, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println("can`t read argument", err)
			continue
		}

		fmt.Println("please enter envelope #2 side b<int>")
		scanner.Scan()
		str = scanner.Text()
		if d, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println("can`t read argument", err)
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
		s = strings.ToLower(strings.Trim(s, ""))
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

func (e Envelope) Fit(e1 Envelope) {
	if e.A > e1.A && e.B > e1.B {
		fmt.Printf("envelope #1: (%v, %v) can accomodate envelope #2: (%v, %v)\n", e.A, e.B, e1.A, e1.B)
	} else {
		fmt.Printf("envelope #1: (%v, %v) can't accomodate envelope #2: (%v, %v)\n", e.A, e.B, e1.A, e1.B)
	}
}
