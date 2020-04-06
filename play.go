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

	fmt.Println("row: ", len(b.state[0]), " col: ", len(b.state))
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter column: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "q" {
			break
		}

		i1, err := strconv.Atoi(text)
		if err == nil && b.moves%2 == 0 {
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
		if err == nil && b.moves%2 == 1 {
			_, err := b.add(2, i1)
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
