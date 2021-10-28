package main

import "testing"

func TestEnvelopeFit(t *testing.T) {
	tests := []struct {
		name string
		e1   Envelope
		e2   Envelope
		want bool
	}{
		{name: "e1.A<e2.A", e1: Envelope{A: 10.0, B: 5.4}, e2: Envelope{A: 10.1, B: 5.0}, want: false},
		{name: "e1.B<e2.B", e1: Envelope{A: 9.0, B: 5.4}, e2: Envelope{A: 8.0, B: 8.1}, want: false},
		{name: "e1.A>e2.B && e1.B > e2.A", e1: Envelope{A: 4.0, B: 11.4}, e2: Envelope{A: 10.2, B: 3.1}, want: true},
		{name: "e1.A>e2.A && e1.A > e2.B", e1: Envelope{A: 14.6, B: 4.4}, e2: Envelope{A: 10.2, B: 3.1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.e1.Fit(tt.e2); got != tt.want {
				t.Errorf("%s: %v.fit(%v) = %t, want %t", tt.name, tt.e1, tt.e2, got, tt.want)
			}
		})
	}
}
