
import (
	"errors"
	"strconv"
	"testing"
)

// func readArgs(args []string) (int, error)
func TestReadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		num  int
		err  error
	}{
		{name: "args!=2", args: []string{"main", "1", "2"}, num: 0, err: errArgsNum},
		{name: "num=NaN", args: []string{"main", "a"}, num: 0, err: strconv.ErrSyntax},
		{name: "ok", args: []string{"main", "100"}, num: 100, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			num, err := readArgs(tt.args)

			if num != tt.num || !errors.Is(err, tt.err) {
				t.Errorf("%s: readArgs(%v) = %d, %v, want %d, %v", tt.name, tt.args, num, err, tt.num, tt.err)
			}
		})
	}
}

// func NumToText(i int, mapper *NumMapper) string {
func TestNumToText(t *testing.T) {
	mapper := initNumberMapper()
	tests := []struct {
		name   string
		i      int
		mapper *NumMapper
		want   string
	}{
		{name: "i=0", i: 0, want: "ноль"},
		{name: "i=-2", i: -2, want: "минус два"},
		{name: "i=1 000 101 418", i: 1000101418, want: "один миллиард сто одна тысяча четыреста восемнадцать"},
		{name: "i=12 112 000", i: 12112000, want: "двенадцать миллионов сто двенадцать тысяч"},
		{name: "i=1 000 000 000", i: 1000000000, want: "один миллиард"},
		{name: "i=14 000 001", i: 14000001, want: "четырнадцать миллионов один"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := NumToText(tt.i, mapper)

			if res != tt.want {
				t.Errorf("%s: NumToText(%d, mapper) = %s, want %s", tt.name, tt.i, res, tt.want)
			}
		})
	}
}
