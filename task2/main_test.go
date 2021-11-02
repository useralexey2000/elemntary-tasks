package main

import (
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestNewEnvelope(t *testing.T) {
	tests := []struct {
		name string
		id   int
		a    float64
		b    float64
		e    *Envelope
		err  error
	}{
		{name: "a<=0", id: 1, a: 0, b: 4, e: nil, err: zeroSideErr},
		{name: "b>=0", id: 1, a: 4, b: 0, e: nil, err: zeroSideErr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env, err := NewEnvelope(tt.id, tt.a, tt.b)
			if err != tt.err || !reflect.DeepEqual(env, tt.e) {
				t.Errorf("%s: NewEnvelope(%d, %f, %f) = %v,%v want %v, %v", tt.name, tt.id, tt.a, tt.b, env, err, tt.e, tt.err)
			}
		})
	}
}

func TestEnvelopeFit(t *testing.T) {
	tests := []struct {
		name string
		e1   *Envelope
		e2   *Envelope
		want bool
	}{
		{name: "e1.A<e2.A", e1: &Envelope{id: 1, A: 10.0, B: 5.4}, e2: &Envelope{id: 2, A: 10.1, B: 5.0}, want: false},
		{name: "e1.B<e2.B", e1: &Envelope{A: 9.0, B: 5.4}, e2: &Envelope{A: 8.0, B: 8.1}, want: false},
		{name: "e1.A>e2.B && e1.B > e2.A", e1: &Envelope{A: 4.0, B: 11.4}, e2: &Envelope{A: 10.2, B: 3.1}, want: true},
		{name: "e1.A>e2.A && e1.A > e2.B", e1: &Envelope{A: 14.6, B: 4.4}, e2: &Envelope{A: 10.2, B: 3.1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.e1.Fit(tt.e2); got != tt.want {
				t.Errorf("%s: %v.fit(%v) = %t, want %t", tt.name, tt.e1, tt.e2, got, tt.want)
			}
		})
	}
}

// other errors need to be checked
func TestReadEnvelope(t *testing.T) {
	// var scanner *bufio.Scanner

	tests := []struct {
		name   string
		reader io.Reader
		id     int
		e      *Envelope
		err    error
	}{
		{name: "canread", reader: strings.NewReader("10.0\n5.4\n"), id: 1, e: &Envelope{id: 1, A: 10.0, B: 5.4}, err: nil},
		{name: "e.A=str", reader: strings.NewReader("a\n5.4\n"), id: 1, e: nil, err: strconv.ErrSyntax},
		{name: "e.B=str", reader: strings.NewReader("4.0\na\n"), id: 1, e: nil, err: strconv.ErrSyntax},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			e1, err := readEnvelope(tt.reader, tt.id)

			if !errors.Is(err, tt.err) || !reflect.DeepEqual(tt.e, e1) {
				t.Errorf("%s: readEnvelope(reader, %d) = %v, %T want: %v, %T", tt.name, tt.id, e1, err, tt.e, tt.err)

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
