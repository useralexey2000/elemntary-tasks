package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minus = "минус"

var errArgsNum = errors.New("not correct number of args")
var errNotPermitedValue = errors.New("not permited value")

func main() {

	i, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	numbers := initNumberMapper()

	text := NumToText(i, numbers)

	fmt.Println(text)
}

func NumToText(i int, numbers map[int]map[int]string) string {

	arr := constructNum(i)
	fmt.Println(arr)

	strarr := numToArrText(arr, numbers)
	fmt.Println(strarr)

	s := numArrTextToText(strarr)
	return s
}

// takes num and returns blocks of ints by category: thousands,mlns,billns
func splitThousand(i int) []int {

	arr := make([]int, 0)

	if i == 0 {
		arr = append(arr, 0)
		return arr
	}

	for i > 0 {
		next := i / 1000
		block := i % 1000
		arr = append([]int{block}, arr...)
		i = next
	}

	return arr
}

// takes num and returns splited hundred by category: ones, tens, hundreds
func splitHundred(i int) []int {

	arr := make([]int, 0)

	var hundred int
	if i > 99 {
		hundred = i / 100
		i = i % 100
	}

	arr = append(arr, hundred)

	var ten int
	if i > 19 {
		ten = i / 10
		i = i % 10
	}

	arr = append(arr, ten)
	arr = append(arr, i)

	return arr
}

// created only for +- sign
type Num struct {
	positive bool
	val      [][]int
}

func constructNum(i int) *Num {
	num := &Num{
		positive: true,
		val:      make([][]int, 0),
	}

	if i < 0 {
		num.positive = false
		i = -1 * i
	}

	thousands := splitThousand(i)
	for _, v := range thousands {

		hundreds := splitHundred(v)
		num.val = append(num.val, hundreds)
	}

	return num
}

func numToArrText(num *Num, m map[int]map[int]string) [][]string {

	arr := make([][]string, 0)
	for i, rank := len(num.val)-1, 0; i >= 0; i, rank = i-1, rank+1 {

		numberBlock := hundredToText(rank, num.val[i], m)
		// skip if 000
		if len(numberBlock) == 0 {
			continue
		}

		// first block doesent have appending rank
		if rank == 0 {
			arr = append([][]string{numberBlock}, arr...)
			continue
		}

		// all values are same after 4 for pos > 0 (10 11 20 40 100 ...thousands / hundreds ...)
		appendingRank := 5
		lastNum := num.val[i][2]
		if lastNum <= 4 && lastNum != 0 {
			appendingRank = lastNum
		}

		//  append rank
		numberBlock = append(numberBlock, m[rank+2][appendingRank])
		arr = append([][]string{numberBlock}, arr...)
	}

	if !num.positive {
		arr = append([][]string{{minus}}, arr...)
	}

	// handle zero
	if len(arr) == 0 {
		arr = append(arr, []string{m[0][0]})
		return arr
	}

	return arr

}

func hundredToText(rank int, arr []int, m map[int]map[int]string) []string {

	block := make([]string, 0)
	for j, pos := 2, 0; j >= 0; j, pos = j-1, pos+1 {

		val := arr[j]
		if val == 0 {
			continue
		}

		str := m[pos][val]
		//  if thousand and it`s 1 || 2
		if rank == 1 && pos == 0 && (val == 1 || val == 2) {
			str = strings.Split(str, "|")[1]
		}

		// if not thousand nad it`s 1 || 2
		if rank != 1 && pos == 0 && (val == 1 || val == 2) {
			str = strings.Split(str, "|")[0]
		}

		block = append([]string{str}, block...)
	}
	return block
}

func numArrTextToText(arr [][]string) string {
	var sb strings.Builder
	sep := " "
	for i, v := range arr {
		s := strings.Join(v, sep)
		if i == len(arr)-1 {
			sb.WriteString(s)
			break
		}
		sb.WriteString(s)
		sb.WriteString(sep)
	}
	return sb.String()
}

func readArgs(args []string) (int, error) {

	if len(args) != 2 {
		return 0, errArgsNum
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, fmt.Errorf("cant read arg num %w", err)
	}

	return num, nil
}

func usage(n string) {
	fmt.Printf("usage: %v number<int>\n", n)
}

func initNumberMapper() map[int]map[int]string {

	ones := map[int]string{
		0:  "ноль",
		1:  "один|одна",
		2:  "два|две",
		3:  "три",
		4:  "четыре",
		5:  "пять",
		6:  "шесть",
		7:  "семь",
		8:  "восемь",
		9:  "девять",
		10: "десять",
		11: "одиннадцать",
		12: "двенадцать",
		13: "тринадцать",
		14: "четырнадцать",
		15: "пятнадцать",
		16: "шестнадцать",
		17: "семнадцать",
		18: "восемнадцать",
		19: "девятнадцать",
	}

	tens := map[int]string{
		2: "двадцать",
		3: "тридцать",
		4: "сорок",
		5: "пятьдесят",
		6: "шестьдесят",
		7: "семыдесят",
		8: "восемьдесят",
		9: "девяносто",
	}

	hundreds := map[int]string{
		1: "сто",
		2: "двести",
		3: "триста",
		4: "четыреста",
		5: "пятьсот",
		6: "шестьсот",
		7: "семысот",
		8: "восемьсот",
		9: "девятьсот",
	}

	thousands := map[int]string{
		1: "тысяча",
		2: "тысячи",
		3: "тысячи",
		4: "тысячи",
		5: "тысяч",
	}

	millions := map[int]string{
		1: "миллион",
		2: "миллиона",
		3: "миллиона",
		4: "миллиона",
		5: "миллионов",
	}

	billions := map[int]string{
		1: "миллиард",
		2: "миллиарда",
		3: "миллиарда",
		4: "миллиарда",
		5: "миллиардов",
	}

	numbers := map[int]map[int]string{
		0: ones,
		1: tens,
		2: hundreds,
		3: thousands,
		4: millions,
		5: billions,
	}

	return numbers
}
