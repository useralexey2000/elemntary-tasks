package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	squares := &TrSquaresHeap{}
	heap.Init(squares)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		//<имя>, <длина стороны>, <длина стороны>, <длина стороны>
		fmt.Println("Please enter triangle parameters: name<string>,sideA<float>,sideB<float>,sideC<float>")
		scanner.Scan()

		str := scanner.Text()
		t, err := stringTriangle(str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Calculate square
		trSq := &TrSquare{
			Name:   t.Name,
			Square: t.Square(),
		}
		// Add to heap
		heap.Push(squares, trSq)

		fmt.Println("Do you want to start again?<yes|y>")
		var s string
		if _, err := fmt.Scanln(&s); err != nil {
			fmt.Println("can`t read argument", err)
			printTriangleSquares(squares)
			break
		}

		s = strings.ToLower(strings.Trim(s, " "))
		if s == "yes" || s == "y" {
			continue
		}

		printTriangleSquares(squares)
		break
	}
}

type Triangle struct {
	Name string
	A    float64
	B    float64
	C    float64
}

// Guron func
func (t *Triangle) Square() float64 {
	p := 0.5 * (t.A + t.B + t.C)
	s := math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	return s
}

func stringTriangle(t string) (*Triangle, error) {
	str := strings.ToLower(t)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\t", "")
	arr := strings.Split(str, ",")

	name := arr[0]

	a, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return nil, fmt.Errorf("Can't parse triangle side A %s", err)
	}

	b, err := strconv.ParseFloat(arr[2], 64)
	if err != nil {
		return nil, fmt.Errorf("Can't parse triangle side B %s", err)
	}

	c, err := strconv.ParseFloat(arr[3], 64)
	if err != nil {
		return nil, fmt.Errorf("Can't parse triangle side C %s", err)
	}

	tr := &Triangle{
		Name: name,
		A:    a,
		B:    b,
		C:    c,
	}
	return tr, nil
}

func printTriangleSquares(sq *TrSquaresHeap) {
	for sq.Len() > 0 {
		i := heap.Pop(sq)
		// Type assertion check
		trSq := i.(*TrSquare)
		fmt.Printf("[%s]: %fm\n", trSq.Name, trSq.Square)
	}
}
