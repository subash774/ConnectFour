package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var b board = board{
		moves: 0,
		state: [][]string{
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		},
	}

	b.show()
	for {
		if b.moves%2 == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter column: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			if text == "q" {
				break
			}

			i1, err := strconv.Atoi(text)
			if err == nil {
				if i1 > 7 {
					fmt.Println("No such column exists, choose a vlaue between 1 and 7")
					continue
				}
				_, err := b.add(1, i1)

				gameOver := b.endGame(1)
				if gameOver {
					b.show()
					fmt.Println("PLAYER 1 WINS!")
					break
				}
				if !err {
					b.moves++
					b.show()
				}
				continue
			}
		}

		if b.moves%2 == 1 {
			p := aiMove("R")
			_, err := b.add(2, p)
			gameOver := b.endGame(2)
			if gameOver {
				b.show()
				fmt.Println("PLAYER 2 WINS")
				break
			}
			if !err {
				b.moves++
				b.show()
			}
			continue

		}

	}
}
