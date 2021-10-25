package main

import (
	"fmt"
	"strings"
)

func main() {
	ones := make(map[int]string)
	ones[0] = "ноль"
	ones[1] = "один|одна"
	ones[2] = "два|две"
	ones[3] = "три"
	ones[4] = "четыре"
	ones[5] = "пять"
	ones[6] = "шесть"
	ones[7] = "семь"
	ones[8] = "восемь"
	ones[9] = "девять"

	teens := make(map[int]string)
	teens[10] = "десять"
	teens[11] = "одинадцать"
	teens[12] = "двенадцать"
	teens[13] = "тринадцать"
	teens[14] = "четырнадцать"
	teens[15] = "пятнадцать"
	teens[16] = "шестнадцать"
	teens[17] = "семнадцать"
	teens[18] = "восемнадцать"
	teens[19] = "девятнадцать"

	tens := make(map[int]string)
	tens[2] = "двадцать"
	tens[3] = "тридцать"
	tens[4] = "сорок"
	tens[5] = "пятьдесят"
	tens[6] = "шестьдесят"
	tens[7] = "семыдесят"
	tens[8] = "восемьдесят"
	tens[9] = "девяносто"

	hundreds := make(map[int]string)
	hundreds[1] = "сто"
	hundreds[2] = "двести"
	hundreds[3] = "триста"
	hundreds[4] = "четыреста"
	hundreds[5] = "пятьсот"
	hundreds[6] = "шестьсот"
	hundreds[7] = "семысот"
	hundreds[8] = "восемьсот"
	hundreds[9] = "девятьсот"

	thousands := make(map[int]string)
	thousands[0] = "тысяч"
	thousands[1] = "тысяча"
	thousands[2] = "тысячи"
	thousands[3] = "тысячи"
	thousands[4] = "тысячи"
	thousands[5] = "тысяч"
	thousands[6] = "тысяч"
	thousands[7] = "тысяч"
	thousands[8] = "тысяч"
	thousands[9] = "тысяч"

	millions := make(map[int]string)
	millions[0] = "миллионов"
	millions[1] = "миллион"
	millions[2] = "миллиона"
	millions[3] = "миллиона"
	millions[4] = "миллиона"
	millions[5] = "миллионов"
	millions[6] = "миллионов"
	millions[7] = "миллионов"
	millions[8] = "миллионов"
	millions[9] = "миллионов"

	billions := make(map[int]string)
	billions[0] = "миллиардов"
	billions[1] = "миллиард"
	billions[2] = "миллиарда"
	billions[3] = "миллиарда"
	billions[4] = "миллиарда"
	billions[5] = "миллиардов"
	billions[6] = "миллиардов"
	billions[7] = "миллиардов"
	billions[8] = "миллиардов"
	billions[9] = "миллиардов"

	numbers := make(map[int]map[int]string)
	numbers[0] = ones
	numbers[1] = teens
	numbers[2] = tens
	numbers[3] = hundreds
	numbers[4] = thousands
	numbers[5] = millions
	numbers[6] = billions

	i := 15001110
	// i := 0
	arr := foo(i)
	nums := bar(arr)
	fmt.Println(arr)
	fmt.Println(nums)
	s := sprint(nums, numbers)
	fmt.Println(s)
}

func foo(i int) []int {
	arr := make([]int, 0)

	if i == 0 {
		arr = append(arr, 0)
		return arr
	}

	for i > 0 {
		j := i / 1000
		k := i % 1000
		arr = append(arr, k)
		i = j
	}
	fmt.Println(arr)
	return arr
}

type num struct {
	val  int
	pos  int
	rank int
}

// ones   // teens   // tens      // hundreds
// (0<i<10) (9<i<20) (19<i<100 ) (99<i<1000)

func bar(arr []int) [][]num {
	nums := make([][]num, 0)

	for i, v := range arr {

		tmp := make([]num, 0)

		if len(arr) == 1 && v == 0 {
			n := num{
				val:  v,
				pos:  0,
				rank: i,
			}
			tmp = append(tmp, n)
			nums = append(nums, tmp)
			return nums
		}

		if v > 99 && v < 1000 {
			val := v / 100
			n := num{
				val:  val,
				pos:  3,
				rank: i,
			}
			tmp = append(tmp, n)
			v = v % 100
		}

		if v > 19 && v < 100 {
			val := v / 10
			n := num{
				val:  val,
				pos:  2,
				rank: i,
			}
			tmp = append(tmp, n)
			v = v % 10
		}

		if v > 9 && v < 20 {
			n := num{
				val:  v,
				pos:  1,
				rank: i,
			}
			tmp = append(tmp, n)
			nums = append(nums, tmp)
			continue
		}
		if v > 0 && v < 10 {
			n := num{
				val:  v,
				pos:  0,
				rank: i,
			}
			tmp = append(tmp, n)
			nums = append(nums, tmp)
			continue
		}

		// can be omited, empty [] arr is returned then
		//  if block is 0 add [0]
		if len(tmp) == 0 && v == 0 {
			n := num{
				val:  v,
				pos:  0,
				rank: i,
			}
			tmp = append(tmp, n)
			nums = append(nums, tmp)
		}

		nums = append(nums, tmp)

	}

	return nums
}

func sprint(nums [][]num, m map[int]map[int]string) string {

	arr := make([][]string, 0)

	for i, v := range nums {

		// if num and num == 0 return 0
		if i == 0 && len(nums) == 1 && v[0].val == 0 {
			return m[0][0]
		}

		// if 000 ignore
		if len(v) > 0 && v[0].val == 0 {
			continue
		}
		// to append from a block
		tmp := make([]string, 0)

		//  n is last number in block | depending on it will be the block named
		var n num
		for _, k := range v {

			n = k
			// if 0 is present in a block ignore
			if n.val == 0 {
				continue
			}

			// get string coresponding to value
			str := m[n.pos][n.val]

			// (+3 added because blocks are devided by hundreds=0, thousands=1 etc...
			// if rank = 1(+3) thousands  and value is 1 || 2 change to second name of <name1>|<name2>
			if n.rank+3 == 4 && n.pos == 0 && (n.val == 1 || n.val == 2) {
				str = strings.Split(str, "|")[1]
			}
			// otherwise change to first name  <name1>|<name2>
			if n.pos == 0 && (n.val == 1 || n.val == 2) {
				str = strings.Split(str, "|")[0]
			}

			tmp = append(tmp, str)
		}
		// dont add rank for blocks < 1000. hundred block dont have additiounal rank
		if n.rank == 0 {
			// if i == 0 {
			arr = append(arr, tmp)
			continue
		}

		// // all values are same after 4 | For all values with rank > 0
		// if n.val > 4 {
		// 	n.val = 5
		// }

		// For all values with rank > 0
		// all values are same after 4 for pos > 1 (10 11 20 40 100 ...thousands / hundreds ...)
		if n.pos >= 1 {
			n.val = 5
		}

		// append rank
		tmp = append(tmp, m[n.rank+3][n.val])
		arr = append(arr, tmp)
	}
	// compile from arr to string
	str := ""
	for i := len(arr) - 1; i >= 0; i-- {
		s := strings.Join(arr[i], " ")
		str = str + s + " "
	}
	return strings.TrimRight(str, " ")
}
