package main

import (
	"bufio"
	"strings"
	"testing"
)

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
		name    string
		scanner *bufio.Scanner
		id      int
		e       *Envelope
		err     error
	}{
		{name: "envelopcorrect", scanner: bufio.NewScanner(strings.NewReader("10.0\n5.4\n")), id: 1, e: &Envelope{id: 1, A: 10.0, B: 5.4}, err: nil},
		{name: "e.A <=0", scanner: bufio.NewScanner(strings.NewReader("0.0\n5.4\n")), id: 1, e: nil, err: zeroSideErr},
		{name: "e.B <=0", scanner: bufio.NewScanner(strings.NewReader("4.0\n0.0\n")), id: 1, e: nil, err: zeroSideErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			e1, err := readEnvelope(tt.scanner, tt.id)
			if err != tt.err && *e1 != *tt.e {
				t.Errorf("%s: readEnvelope(scanner, %d) = %v want: %v", tt.name, tt.id, tt.e, e1)

			}
		})
	}
}

func TestAskRepeat(t *testing.T) {
	tests := []struct {
		name    string
		scanner *bufio.Scanner
		want    bool
	}{
		{name: "yeswithspace", scanner: bufio.NewScanner(strings.NewReader("yes  ")), want: true},
		{name: "ywithspace", scanner: bufio.NewScanner(strings.NewReader("  y  ")), want: true},
		{name: "Yuppercase", scanner: bufio.NewScanner(strings.NewReader(" Y ")), want: true},
		{name: "no", scanner: bufio.NewScanner(strings.NewReader("no")), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := askRepeat(tt.scanner)
			if b != tt.want {
				t.Errorf("%s: askRepeat(scanner) = %t want: %t", tt.name, b, tt.want)

			}
		})
	}
}
