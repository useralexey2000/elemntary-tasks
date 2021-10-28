package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var zeroSideErr = errors.New("side cant be <=0")

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		e1, err := readEnvelope(scanner, 1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		e2, err := readEnvelope(scanner, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		b := e1.Fit(e2)
		print(e1, e2, b)

		if askRepeat(scanner) {
			continue
		}

		break
	}
}

type Envelope struct {
	id int
	A  float64
	B  float64
}

func (e *Envelope) Fit(e1 *Envelope) bool {
	if e.A > e1.A && e.B > e1.B || e.A > e1.B && e.B > e1.A {
		return true
	}
	return false
}

func print(e1, e2 *Envelope, b bool) {
	if b {
		fmt.Printf("envelope #%d: (%f, %f) can accomodate envelope #%d: (%f, %f)\n", e1.id, e1.A, e1.B, e2.id, e2.A, e2.B)
		return
	}
	fmt.Printf("envelope #%d: (%f, %f) can't accomodate envelope #%d: (%f, %f)\n", e1.id, e1.A, e1.B, e2.id, e2.A, e2.B)
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

func readEnvelope(scanner *bufio.Scanner, id int) (*Envelope, error) {

	fmt.Printf("please enter envelope %d side a<float>\n", id)

	scanner.Scan()
	str := scanner.Text()
	a, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return nil, fmt.Errorf("can`t read argument %v", err)
	}

	if a <= 0 {
		return nil, zeroSideErr
	}

	fmt.Printf("please enter envelope %d side b<float>\n", id)

	scanner.Scan()
	str = scanner.Text()
	b, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return nil, fmt.Errorf("can`t read argument %v", err)
	}

	if b <= 0 {
		return nil, zeroSideErr
	}

	return &Envelope{id: id, A: a, B: b}, nil
}

func askRepeat(scanner *bufio.Scanner) bool {
	fmt.Println("Do you want to start again?<yes|y>")

	scanner.Scan()
	str := scanner.Text()
	str = strings.ToLower(strings.Trim(str, " "))

	if str == "yes" || str == "y" {
		return true
	}

	return false
}
