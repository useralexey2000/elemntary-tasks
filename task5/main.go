package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var errArgsNum = errors.New("not correct number of args")
var errExeedsPermitedValue = errors.New("number exeeds permited value")

func main() {

	numbers := initNumberMapper()

	i := 3211810111

	arr := splitNumber(i)
	fmt.Println(arr)

	strarr := sprintArrNum(arr, numbers)

	fmt.Println(strarr)

	s := blocksToString(strarr)
	fmt.Println(s)
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

func splitNumber(i int) [][]int {
	arr := make([][]int, 0)

	thousands := splitThousand(i)

	for _, v := range thousands {
		hundreds := splitHundred(v)
		arr = append(arr, hundreds)
	}

	return arr
}

func sprintArrNum(nums [][]int, m map[int]map[int]string) [][]string {

	arr := make([][]string, 0)

	for i, rank := len(nums)-1, 0; i >= 0; i, rank = i-1, rank+1 {

		numberBlock := make([]string, 0)

		// all values are same after 4 for pos > 0 (10 11 20 40 100 ...thousands / hundreds ...)
		appendingRank := 5
		for j, pos := 2, 0; j >= 0; j, pos = j-1, pos+1 {

			val := nums[i][j]

			if val == 0 {
				continue
			}

			if j == 2 && val <= 4 {
				appendingRank = val
			}

			str := m[pos][val]
			//  if thousand and it`s 1 || 2
			if i == len(nums)-2 && pos == 0 && (val == 1 || val == 2) {
				str = strings.Split(str, "|")[1]
			}

			// if not thousand nad it`s 1 || 2
			if pos == 0 && (val == 1 || val == 2) {
				str = strings.Split(str, "|")[0]
			}

			numberBlock = append([]string{str}, numberBlock...)

		}
		// skip if 000
		if len(numberBlock) == 0 {
			continue
		}

		// first block doesent have appending rank
		if i == len(nums)-1 {
			arr = append([][]string{numberBlock}, arr...)
			continue
		}

		//  append rank
		numberBlock = append(numberBlock, m[rank+2][appendingRank])

		arr = append([][]string{numberBlock}, arr...)

	}

	// handle zero
	if len(arr) == 0 {
		arr = append(arr, []string{m[0][0]})
		return arr
	}

	return arr

}

func blocksToString(arr [][]string) string {
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
		11: "одинадцать",
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
