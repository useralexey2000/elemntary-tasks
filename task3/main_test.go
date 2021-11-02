package main

import (
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// func check(a, b, c float64) bool {
func TestCheck(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c float64
		ok      bool
	}{
		{name: "correct", a: 1.1, b: 2.0, c: 3.2, ok: true},
		{name: "a<0", a: -1.1, b: 2.0, c: 3.2, ok: false},
		{name: "b==0", a: 1.1, b: 0, c: 3.2, ok: false},
		{name: "c<0", a: 1.1, b: 0.1, c: -3.2, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ok := check(tt.a, tt.b, tt.c)

			if tt.ok != ok {
				t.Errorf("%s: check(%f, %f, %f) = %t want: %t", tt.name, tt.a, tt.b, tt.c, ok, tt.ok)

			}
		})
	}
}

// func NewTriangle(name string, a, b, c float64) (*Triangle, error) {
func TestNewTriangle(t *testing.T) {
	tests := []struct {
		name    string
		trname  string
		a, b, c float64
		tr      *Triangle
		err     error
	}{
		{name: "correct", trname: "triangle1", a: 5.1, b: 6.5, c: 4.0, tr: &Triangle{Name: "triangle1", A: 5.1, B: 6.5, C: 4.0}, err: nil},
		{name: "sideA=0", trname: "triangle1", a: 0, b: 6.5, c: 4.0, tr: nil, err: notTriangleErr},
		{name: "sideB=0", trname: "triangle1", a: 5.1, b: 0, c: 4.0, tr: nil, err: notTriangleErr},
		{name: "sideC=-1", trname: "triangle1", a: 5.1, b: 6.5, c: -1, tr: nil, err: notTriangleErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tr, err := NewTriangle(tt.trname, tt.a, tt.b, tt.c)

			if !errors.Is(err, tt.err) || !reflect.DeepEqual(tt.tr, tr) {
				t.Errorf("%s: NewTriangle(%s, %f, %f, %f) = %v, %v want: %v, %v", tt.name, tt.trname, tt.a, tt.b, tt.c, tr, err, tt.tr, tt.err)

			}
		})
	}
}

// func (t *Triangle) Square() float64 {

func TestSquare(t *testing.T) {
	tests := []struct {
		name string
		tr   *Triangle
		want float64
	}{
		{name: "correct", tr: &Triangle{Name: "triangle1", A: 5.1, B: 6.5, C: 4.0}, want: 10.199823527885176},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tr.Square()

			if got != tt.want {
				t.Errorf("%s: tr.square(*tr) = %f want: %f", tt.name, got, tt.want)

			}
		})
	}
}

// func readTriangle(r io.Reader) (*Triangle, error) {
func TestReadTriangle(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		tr     *Triangle
		err    error
	}{
		{name: "canread", reader: strings.NewReader("triangle1,10.0,5.4,10.1"), tr: &Triangle{Name: "triangle1", A: 10.0, B: 5.4, C: 10.1}, err: nil},
		{name: "tr.A=str", reader: strings.NewReader("triangle1,a,5.4,10.1"), tr: nil, err: strconv.ErrSyntax},
		{name: "tr.B=str", reader: strings.NewReader("triangle1,10.0,b,10.1"), tr: nil, err: strconv.ErrSyntax},
		{name: "tr.C=str", reader: strings.NewReader("triangle1,10.0,5.4,c"), tr: nil, err: strconv.ErrSyntax},
		{name: "args!=4", reader: strings.NewReader("triangle1,10.0"), tr: nil, err: notCorrectArgsLenErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tr, err := readTriangle(tt.reader)

			if !errors.Is(err, tt.err) || !reflect.DeepEqual(tt.tr, tr) {
				t.Errorf("%s: readTriangle(reader) = %v, %v want: %v, %v", tt.name, tr, err, tt.tr, tt.err)

			}
		})
	}
}

func TestAskRepeat(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   bool
	}{
		{name: "yeswithspace", reader: strings.NewReader("yes  "), want: true},
		{name: "ywithspace", reader: strings.NewReader("  y  "), want: true},
		{name: "Yuppercase", reader: strings.NewReader(" Y "), want: true},
		{name: "no", reader: strings.NewReader("no"), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := askRepeat(tt.reader)
			if b != tt.want {
				t.Errorf("%s: askRepeat(scanner) = %t want: %t", tt.name, b, tt.want)

			}
		})
	}
}
