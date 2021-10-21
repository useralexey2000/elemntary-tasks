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

	// teens := make(map[string]string)
	// tens["0"] = "десять"
	// tens["1"] = "одинадцать"
	// tens["2"] = "двенадцать"
	// tens["3"] = "тринадцать"
	// tens["4"] = "четырнадцать"
	// tens["5"] = "пятнадцать"
	// tens["6"] = "шестнадцать"
	// tens["7"] = "семнадцать"
	// tens["8"] = "восемнадцать"
	// tens["9"] = "девятнадцать"

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
	tens["2"] = "двадцать"
	tens["3"] = "тридцать"
	tens["4"] = "сорок"
	tens["5"] = "пятьдесят"
	tens["6"] = "шестьдесят"
	tens["7"] = "семыдесят"
	tens["8"] = "восемьдесят"
	tens["9"] = "девяносто"

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
	thousands["0"] = "тысяч"
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
	millions["0"] = "миллионов"
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
	billions["0"] = "миллиардов"
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
	numbers[0] = ones
	numbers[1] = tens
	numbers[2] = hundreds
	numbers[3] = thousands
	numbers[4] = millions
	numbers[5] = billions

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

	str := "41200"

	arr := partNumString(str)

	// res := ""
	// for i, v := range arr {
	// 	name := string(v[len(v)-1])
	// 	k := strings.Split(v, "")

	// 	idx := len(v) - 1
	// 	part := ""
	// 	for j := 0; j <= len(k)-1; j++ {

	// 		if k[j] == "0" {
	// 			idx--
	// 			continue
	// 		}
	// 		if len(k) == 2 && k[0] == "1" {
	// 			part = numbers[i][strings.Join(k, "")]
	// 			break
	// 		}

	// 		if j == 1 && k[j] == "1" {
	// 			pp := numbers[idx][k[j]+k[j+1]]
	// 			part = part + " " + pp
	// 			break
	// 		}

	// 		p := numbers[idx][k[j]]

	// 		idx--
	// 		// part = part + " " + p
	// 		part = part + " " + p
	// 	}

	// 	identifier := i + 2
	// 	// if thousanda

	// 	if i < 1 {
	// 		if part == "" {
	// 			res = part + res
	// 			continue
	// 		}
	// 		res = part + numbers[len(arr)-1][string(part[0])] + res
	// 		continue
	// 	}
	// 	if part == "" {
	// 		res = part + res
	// 		continue
	// 	}
	// 	rank := numbers[identifier][name]

	// 	// s := string(part[len(part)-1])
	// 	pt := strings.Trim(part, " ")
	// 	if pt == numbers[0]["1"] && identifier == 3 {
	// 		part = strings.Split(part, "|")[1]
	// 	}

	// 	if part == numbers[0]["2"] && identifier == 3 {
	// 		part = strings.Split(part, "|")[1]
	// 	}

	// 	//add identifier of 3 pair len(arr)+1
	// 	res = part + " " + rank + numbers[len(arr)-1][string(part[0])] + res
	// }

	// fmt.Println(res)
	fmt.Println(arr)
}

func usage(s string) {
	fmt.Printf("usage: %s number<int>\n", s)
}

// {"100", 50, 0, 200, 0, 1}
// 11 200
//

func partNumString(str string) []string {
	arr := make([]string, 0)
	count := 0
	// devide by category
	for i := len(str); i > 3; i -= 3 {
		arr = append(arr, string(str[i-3:i]))
		count++
	}
	arr = append(arr, str[0:len(str)-count*3])
	// fmt.Println(arr)
	return arr
}

func printNumStringPart(nums map[int]map[string]string, str string, rank int) []string {
	// var name string

	ss := strings.Split(str, "")
	arr := make([]string, 0)

	idx := 2
	for i := 0; i < len(ss); i++ {
		if ss[i] == "0" {
			idx--
			continue

		}
		if len(ss)-i == 2 && ss[i] == "1" {
			s := ss[i] + ss[i+1]
			fmt.Println(s)
			arr = append(arr, nums[1][s])
			return arr
		}
		arr = append(arr, nums[idx][ss[i]])
		idx--
	}
	fmt.Println(arr)

	return arr
}

// 1 015
// 14
