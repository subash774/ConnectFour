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
			if !err {
				b.moves++
				b.show()
			}
			continue
		}
		if err == nil && b.moves%2 == 1 {
			_, err := b.add(2, i1)
			if !err {
				b.moves++
				b.show()
			}
			continue
		}

	}
}
