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

	turn := "X"
	var x, y int
	for {
		// Show the board to help player decide their move
		// TODO: try to clear the terminal between turns for a cleaner experience
		fmt.Printf("\n%s", b.Display())

		// Get a valid move from a player
	turnLoop:
		for {
			fmt.Printf("It's %q's turn. Enter move in format `x y`: ", turn)
			fmt.Scan(&x, &y)
			// fmt.Printf("your move was `%d %d`", x, y) // use logs and make this
			if b.IsValidMove(x, y) {
				break turnLoop
			}
		}

		b.RegisterMove(turn, x, y)

		// Check whether the last move resulted in a winner
		if b.HasWinner() {
			fmt.Printf("Player %q wins!", turn)
			break
		}

		// If the last move didn't result in a winner, it's the next player's turn
		turn = switchTurn(turn)
	}
}

func switchTurn(turn string) string {
	if turn == "X" {
		return "O"
	}

	return "X"
}
