package main

import (
	"fmt"

	"tic-tac-go/board"
	"tic-tac-go/utils"
)

func main() {
	// Display splash screen
	fmt.Printf(utils.Splash())

	b := board.Board{}
	b.Initialise()

	fmt.Printf(b.Display())
}
