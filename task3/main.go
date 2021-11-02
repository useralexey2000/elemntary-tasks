package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	notCorrectArgsLenErr = errors.New("not correct args length")
	notTriangleErr       = errors.New("not a triangle")
)

func main() {
	squares := &TrSquaresHeap{}
	heap.Init(squares)

	for {
		//<имя>, <длина стороны>, <длина стороны>, <длина стороны>
		fmt.Println("Please enter triangle parameters: name<string>,sideA<float>,sideB<float>,sideC<float>")

		t, err := readTriangle(os.Stdin)
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

		if askRepeat(os.Stdin) {
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

func NewTriangle(name string, a, b, c float64) (*Triangle, error) {
	if !check(a, b, c) {
		return nil, notTriangleErr
	}
	tr := &Triangle{
		Name: name,
		A:    a,
		B:    b,
		C:    c,
	}
	return tr, nil
}

func check(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return true
}

// Guron func
func (t *Triangle) Square() float64 {
	p := 0.5 * (t.A + t.B + t.C)
	s := math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	return s
}

func printTriangleSquares(sq *TrSquaresHeap) {
	for sq.Len() > 0 {
		i := heap.Pop(sq)
		// Type assertion check
		trSq := i.(*TrSquare)
		fmt.Printf("[%s]: %fm\n", trSq.Name, trSq.Square)
	}
}

func readTriangle(r io.Reader) (*Triangle, error) {
	// at least one number should be specified

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	str := scanner.Text()

	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\t", "")
	arr := strings.Split(str, ",")

	if len(arr) != 4 {
		return nil, notCorrectArgsLenErr
	}

	name := arr[0]

	a, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse triangle side A %w", err)
	}

	b, err := strconv.ParseFloat(arr[2], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse triangle side B %w", err)
	}

	c, err := strconv.ParseFloat(arr[3], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse triangle side C %w", err)
	}

	return NewTriangle(name, a, b, c)

}

func askRepeat(r io.Reader) bool {
	fmt.Println("Do you want to start again?<yes|y>")

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	str := scanner.Text()
	str = strings.ToLower(strings.Trim(str, " "))

	if str == "yes" || str == "y" {
		return true
	}

	return false
}
