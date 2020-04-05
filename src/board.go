package main

import (
	"fmt"
)

type board struct {
	moves int
	state [][]string
}

func (b board) add(p int, pos int) board {

	if b.state[0][pos-1] != " | " {
		fmt.Println("Can't add it here")
		return b
	}

	tile := ""
	if p == 1 {
		tile = "X"
	} else {
		tile = "O"
	}

	idx := 1
	for {
		if b.state[len(b.state)-idx][pos-1] == " | " {
			b.state[len(b.state)-idx][pos-1] = tile
			break
		} else {
			idx++
		}
	}
	return b
}

func (b board) show() {
	for _, element := range b.state {
		fmt.Println(element)
	}
}

func main() {
	b := board{
		moves: 0,
		state: [][]string{
			[]string{" | ", " | ", " | ", " | ", " | ", " | ", " | "},
			[]string{" | ", " | ", " | ", " | ", " | ", " | ", " | "},
			[]string{" | ", " | ", " | ", " | ", " | ", " | ", " | "},
			[]string{" | ", " | ", " | ", " | ", " | ", " | ", " | "},
			[]string{" | ", " | ", " | ", " | ", " | ", " | ", " | "},
			[]string{" | ", " | ", " | ", " | ", " | ", " | ", " | "},
		},
	}

	b.add(1, 2)
	b.add(1, 2)
	b.add(2, 2)
	b.add(1, 2)
	b.add(1, 2)
	b.add(2, 2)
	b.add(1, 2)

	b.show()
}
