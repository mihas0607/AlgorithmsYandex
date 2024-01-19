package main

import "fmt"

func main() {
	var (
		input string
		ans string
		anscnt int
	)
	fmt.Scan(&input)
	for i:=0; i<len(input); i++{
		nowcnt := 0
		for j:=0; j<len(input); j++ {
			if input[i] == input[j] {
				nowcnt ++
			}
		}
		if nowcnt > anscnt {
			ans = string (input[i])
			anscnt = nowcnt
		}
	}
	fmt.Println(ans)
}
