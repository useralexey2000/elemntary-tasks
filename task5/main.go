package main

import (
	"fmt"
	"strings"
)

func main() {
	ones := make(map[string]string)
	ones["0"] = "ноль"
	ones["1"] = "один|одна"
	ones["2"] = "два|две"
	ones["3"] = "три"
	ones["4"] = "четыре"
	ones["5"] = "пять"
	ones["6"] = "шесть"
	ones["7"] = "семь"
	ones["8"] = "восемь"
	ones["9"] = "девять"

	tens := make(map[string]string)
	tens["10"] = "десять"
	tens["11"] = "одинадцать"
	tens["12"] = "двенадцать"
	tens["13"] = "тринадцать"
	tens["14"] = "четырнадцать"
	tens["15"] = "пятнадцать"
	tens["16"] = "шестнадцать"
	tens["17"] = "семнадцать"
	tens["18"] = "восемнадцать"
	tens["19"] = "девятнадцать"
	tens["20"] = "двадцать"
	tens["30"] = "тридцать"
	tens["40"] = "сорок"
	tens["50"] = "пятьдесят"
	tens["60"] = "шестьдесят"
	tens["70"] = "семыдесят"
	tens["80"] = "восемьдесят"
	tens["90"] = "девяносто"

	// tens := make(map[string]string)
	// tens["1"] = "десять"
	// tens["2"] = "двадцать"
	// tens["3"] = "тридцать"
	// tens["4"] = "сорок"
	// tens["5"] = "пятьдесят"
	// tens["6"] = "шестьдесят"
	// tens["7"] = "семыдесят"
	// tens["8"] = "восемьдесят"
	// tens["9"] = "девяносто"

	hundreds := make(map[string]string)
	hundreds["1"] = "сто"
	hundreds["2"] = "двести"
	hundreds["3"] = "триста"
	hundreds["4"] = "четыреста"
	hundreds["5"] = "пятьсот"
	hundreds["6"] = "шестьсот"
	hundreds["7"] = "семысот"
	hundreds["8"] = "восемьсот"
	hundreds["9"] = "девятьсот"

	thousands := make(map[string]string)
	thousands["1"] = "тысяча"
	thousands["2"] = "тысячи"
	thousands["3"] = "тысячи"
	thousands["4"] = "тысячи"
	thousands["5"] = "тысяч"
	thousands["6"] = "тысяч"
	thousands["7"] = "тысяч"
	thousands["8"] = "тысяч"
	thousands["9"] = "тысяч"

	millions := make(map[string]string)
	millions["1"] = "миллион"
	millions["2"] = "миллиона"
	millions["3"] = "миллиона"
	millions["4"] = "миллиона"
	millions["5"] = "миллионов"
	millions["6"] = "миллионов"
	millions["7"] = "миллионов"
	millions["8"] = "миллионов"
	millions["9"] = "миллионов"

	billions := make(map[string]string)
	billions["1"] = "миллиард"
	billions["2"] = "миллиарда"
	billions["3"] = "миллиарда"
	billions["4"] = "миллиарда"
	billions["5"] = "миллиардов"
	billions["6"] = "миллиардов"
	billions["7"] = "миллиардов"
	billions["8"] = "миллиардов"
	billions["9"] = "миллиардов"

	numbers := make(map[int]map[string]string)
	numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5] = ones, tens, hundreds, thousands, millions, billions

	// args := os.Args
	// if len(args) < 2 {
	// 	fmt.Println("Please enter number")
	// 	usage(os.Args[0])
	// 	os.Exit(1)
	// }

	// // 900 000 000 000
	// // m t
	// str := os.Args[1]

	// if _, err := strconv.Atoi(str); err != nil {
	// 	fmt.Println("Please enter correct number", err)
	// 	usage(args[0])
	// 	os.Exit(1)
	// }

	str := "15201"
	arr := make([]string, 0)
	count := 0
	// devide by category
	if len(str) > 3 {
		for i := len(str); i > 3; i -= 3 {
			arr = append(arr, string(str[i-3:i]))
			count++
		}
		arr = append(arr, str[0:len(str)-count*3])
		fmt.Println(arr)
	}

	res := ""
	for i, v := range arr {
		// for _, k := range strings.Split(v, "") {
		k := strings.Split(v, "")

		if len(k) == 2 && k[0] == "1" {
			res = numbers[i][strings.Join(k, "")] + res
			continue
		}
		idx := 0
		for j := len(k) - 1; j >= 0; j-- {

			if k[j] == "0" {
				idx++
				continue
			}

			res = numbers[idx][k[j]] + res
			idx++
			fmt.Println(res)
		}
	}

	fmt.Println(res)
}

func usage(s string) {
	fmt.Printf("usage: %s number<int>\n", s)
}

// {"100", 50, 0, 200, 0, 1}
