package main

import (
	"fmt"
)

var sum int

func main() {
	// var sl = []int{1, 2, 3, 4, 5, 6}

	// loop:
	// 	for i := 0; i < len(sl); i++ {
	// 		if sl[i]%2 == 0 {
	// 			continue loop
	// 		}
	// 		if sum > 12 {
	// 			break
	// 		}
	// 		sum += sl[i]
	// 	}

	// fmt.Println("sum is", sum)

	labelForLoop(53)
}

func labelForLoop(num int) {
	var sl = [][]int{
		{1, 2, 3, 4, 5},
		{2, 4, 5, 6, 10},
		{22, 33, 53, 63, 130},
	}
outline:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == num {
				fmt.Printf("find the number %d at position %d, %d\n", num, i, j)
				continue outline
			}
			fmt.Printf("find the number %d at j position %d,\n", num, j)
		}
		fmt.Printf("find the number %d at i position %d,\n", num, i)
	}
}
