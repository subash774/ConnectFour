package main

import (
	"fmt"
)

type board struct {
	moves int
	state [][]string
}

func (b board) getTile(p int) string {
	tile := ""
	if p == 1 {
		tile = "| X |"
	} else {
		tile = "| O |"
	}
	return tile
}

func (b board) endGame(p int) bool {
	rows := 7
	cols := 6
	state := b.state
	tile := b.getTile(p)

	//  horizontal check
	for j := 0; j < rows-3; j++ {
		for i := 0; i < cols; i++ {
			// fmt.Println(i, j)
			if state[i][j] == tile && state[i][j+1] == tile && state[i][j+2] == tile && state[i][j+3] == tile {
				return true
			}
		}
	}

	//  vertical check
	for i := 0; i < cols-3; i++ {
		for j := 0; j < rows; j++ {
			// fmt.Println(i, j)
			if state[i][j] == tile && state[i+1][j] == tile && state[i+2][j] == tile && state[i+3][j] == tile {
				return true
			}
		}
	}

	// +ve diag
	for i := 3; i < cols; i++ {
		for j := 0; j < rows-3; j++ {
			if state[i][j] == tile && state[i-1][j+1] == tile && state[i-2][j+2] == tile && state[i-3][j+3] == tile {
				return true
			}
		}
	}

	// -ve diag
	for i := 0; i < cols-3; i++ {
		for j := 3; j < rows; j++ {
			if state[i][j] == tile && state[i-1][j-1] == tile && state[i-2][j-2] == tile && state[i-3][j-3] == tile {
				return true
			}
		}
	}

	return false

}

func (b board) add(p int, pos int) (board, bool) {

	if b.state[0][pos-1] != "| - |" {
		fmt.Println("Can't add it here")
		return b, true
	}

	tile := b.getTile(p)

	idx := 1
	for {
		if b.state[len(b.state)-idx][pos-1] == "| - |" {
			b.state[len(b.state)-idx][pos-1] = tile
			break
		} else {
			idx++
		}
	}
	return b, false
}

func (b board) show() {
	for _, element := range b.state {
		fmt.Println(element)
	}
}
