package main

import (
	"fmt"
	"math/rand"
)

func aiMove(difficulty string) int {
	if difficulty == "R" {
		fmt.Println(" ... ")
		v := rand.Intn(7-1) + 1
		return v
	}
	return -1
}

// TODO: implement minimax search

// func aiMove(difficulty int) int {
// 	depth := 0
// 	switch difficulty {
// 	case 1:
// 		depth = 3
// 	case 2:
// 		depth = 5
// 	case 3:
// 		depth = 7
// 	case 4:
// 		depth = 10
// 	default:
// 		depth = 3

// 	}

// }
