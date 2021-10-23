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

	str := "1001104000"

	arr := partNumString(str)
	// fmt.Println(arr)
	arr = bar(arr, numbers)

	fmt.Println(arr)
}

type number struct {
	str  string
	rank int
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
		s := string(str[i-3 : i])
		// can be leading zeros
		s = strings.TrimLeft(s, "0")
		if s == "" {
			s = "0"
		}
		arr = append(arr, s)
		count++
	}
	arr = append(arr, str[0:len(str)-count*3])
	fmt.Println(arr, len(arr))
	return arr
}

func bar(arr []string, nums map[int]map[string]string) []string {

	sl := make([]string, 0)
	// tmp := make([]string, 0)
	// tmp := make([]number, 0)

	for _, v := range arr {
		// rank := i + 3

		ss := strings.Split(v, "")

		if len(ss) == 1 {
			num := nums[0][ss[0]]
			sl = append(sl, num)
			continue
		}

		// not complex - teens
		if len(ss) == 2 && ss[0] == "1" {
			num := ss[0] + ss[1]
			num = nums[1][num]
			// num = num + " " + nums[i+1][arr[i]]

			// numb := number{
			// 	str: nums[1][n],
			// 	rank:
			// }
			sl = append(sl, num)
			continue
		}

		tmp := make([]string, 0)
		// complex tens and hundreds
		for j, k := 0, len(ss)-1; k >= 0; j, k = j+1, k-1 {
			s := ss[k]
			if s == "0" && k >= 1 {
				continue
			}

			num := nums[j][s]
			tmp = append(tmp, num)
		}
		str := strings.Join(tmp, " ")
		sl = append(sl, str)
	}

	fmt.Println(sl)
	res := make([]string, 0)
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	if arr[i] == "0" {
	// 		tmp[i] = ""
	// 		continue
	// 	}

	// 	if len(arr[i]) == 1 {
	// 		s := nums[i+2][arr[i]]
	// 		res = append(res, tmp[i]+" "+s)
	// 		continue
	// 	}

	// 	if len(arr[i]) == 2 {
	// 		s := nums[i+1][arr[i]]
	// 		res = append(res, tmp[i]+" "+s)
	// 		continue
	// 	}
	// 	// if len(arr[i]) == 3 {

	// 	// }
	// 	res = append(res, tmp[i])
	// }

	return res
}

// 1 015
// 14
