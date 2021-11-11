package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minus = "минус"
const sep = " "

var errArgsNum = errors.New("not correct number of args")

func main() {

	i, err := readArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		usage(os.Args[0])
		os.Exit(1)
	}

	mapper := initNumberMapper()

	text := NumToText(i, mapper)

	fmt.Println(text)
}

func NumToText(i int, mapper *NumMapper) string {

	arr := constructNum(i)
	// fmt.Println(arr)

	strarr := numToArrText(arr, mapper)
	// fmt.Println(strarr)

	s := numArrTextToText(strarr)
	return s
}

type Num struct {
	positive bool
	block    []*NumBlock
}

// constructs blocks of numbers by rank
// true [[0,0,1], [0,0,0], [1,0,18]] = 1 000 118
func constructNum(i int) *Num {
	positive := true
	if i < 0 {
		positive = false
		i = -1 * i
	}

	thousands := splitThousand(i)

	num := &Num{
		positive: positive,
		block:    make([]*NumBlock, 0),
	}
	rank := len(thousands) - 1
	for _, v := range thousands {

		block := splitHundred(rank, v)
		num.block = append(num.block, block)
		rank--
	}

	return num
}

// takes num and returns blocks of ints by category: thousands,mlns,billns
// [1, 0, 124] = 1 000 124
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
	fmt.Println(arr)
	return arr
}

type NumBlock struct {
	rank int
	val  [3]int
}

// takes num  & rank and returns splited hundred by category: ones, tens, hundreds
// [3, 0, 18] = 318
// [1, 2, 4] = 124
func splitHundred(rank, i int) *NumBlock {

	var hundred int
	if i > 99 {
		hundred = i / 100
		i = i % 100
	}

	var ten int
	if i > 19 {
		ten = i / 10
		i = i % 10
	}

	one := i

	return &NumBlock{
		rank: rank,
		val:  [3]int{hundred, ten, one},
	}
}

// takes num and returns array text number
func numToArrText(num *Num, mapper *NumMapper) [][]string {

	arr := make([][]string, 0)

	for _, block := range num.block {

		textBlock := numBlockToArrText(block, mapper)

		if textBlock != nil {
			arr = append(arr, textBlock)
		}

	}

	if len(arr) == 0 {
		arr = append(arr, []string{mapper.number[0][0]})
		return arr
	}

	if !num.positive {
		arr = append([][]string{{minus}}, arr...)
	}

	return arr
}

//  if thousand pos 0 and it`s 1 || 2 treats differently
func numBlockToArrText(block *NumBlock, mapper *NumMapper) []string {

	textBlock := make([]string, 0)

	pos := 2
	for _, val := range block.val {

		if val == 0 {
			pos--
			continue
		}

		str := mapper.number[pos][val]
		if block.rank == 1 && pos == 0 && (val == 1 || val == 2) {
			str = strings.Split(str, "|")[1]
		}

		if block.rank != 1 && pos == 0 && (val == 1 || val == 2) {
			str = strings.Split(str, "|")[0]
		}

		textBlock = append(textBlock, str)
		pos--
	}
	if len(textBlock) == 0 {
		return nil
	}

	if block.rank == 0 {
		return textBlock
	}

	rankVal := 5
	lastNum := block.val[2]
	if lastNum <= 4 && lastNum != 0 {
		rankVal = lastNum
	}

	textBlock = append(textBlock, mapper.rank[block.rank][rankVal])

	return textBlock
}

func numArrTextToText(arr [][]string) string {
	var sb strings.Builder
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

type NumMapper struct {
	number map[int]map[int]string
	rank   map[int]map[int]string
}

func initNumberMapper() *NumMapper {

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
	}
	ranks := map[int]map[int]string{
		1: thousands,
		2: millions,
		3: billions,
	}
	mapper := &NumMapper{
		number: numbers,
		rank:   ranks,
	}

	return mapper
}
