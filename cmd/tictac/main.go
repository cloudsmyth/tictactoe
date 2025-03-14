package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cloudsmyth/tictactoe/internal/game"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("welcome to tic-tac-toe!")

	play := true

	for play {
		play = game.PlayGame(scanner)
	}

	fmt.Println("Thanks for playing! =]")
}
