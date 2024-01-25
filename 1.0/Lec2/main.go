package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// func EvenOrOdd(number int) string {
// 	if number%2 == 0 {
// 		return "Even"
// 	}
// 	return "Odd"
// }

// func GetCount(str string) (count int) {
// 	// Enter solution here
// 	sliceOfVowels := []rune{'a', 'e', 'i', 'o', 'u'}
// 	for _, val1 := range str {
// 		for _, val2 := range sliceOfVowels {
// 			if val2 == val1 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func Disemvowel(comment string) string {

// 	for _, c := range comment {
// 		switch c {
// 		case 'a', 'e', 'i', 'o', 'u':
// 			comment = strings.ReplaceAll(comment, string(c), "")
// 		}
// 	}
// 	return comment
// }

func HighAndLow(in string) string {
	a := strings.Fields(in)
	sliceNum := []int{}
	for _, v := range a {
		b, _ := strconv.Atoi(string(v))
		sliceNum = append(sliceNum, b)
	}
	sort.Ints(sliceNum)
	fmt.Println(sliceNum)
	newStr := fmt.Sprintf("%d %d", sliceNum[len(sliceNum)-1], sliceNum[0])
	return newStr
}

// func main() {
// 	var inputStr string
// 	// fmt.Scan(&inputStr)

// 	input := bufio.NewScanner(os.Stdin)
// 	if input.Scan() {
// 		inputStr = input.Text()
// 	}
// 	fmt.Println(HighAndLow(inputStr))

// }
