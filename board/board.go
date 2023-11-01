package board

import (
	"bytes"
	"fmt"
)

const (
	Size = 3

	ZeroVal = 0 // This value should not be changed
	P1Val   = 1
	P2Val   = -1
)

var (
	displayMap = map[int]string{
		ZeroVal: " ",
		P1Val:   "X",
		P2Val:   "O",
	}
)

type Board struct {
	Tiles [Size][Size]int
}

// Create a 3x3 two-dimensional array and fill it with the zero value
func (b Board) InitialiseBoard() {
	b.Tiles = [Size][Size]int{}
}

// Determine whether a proposed move is valid by checking for collisions
// with non-zero values
func (b Board) IsValidMove(x, y int) bool {
	if b.Tiles[x][y] != ZeroVal {
		return false
	}

	return true
}

// ...
func (b Board) HasWinner() bool {
	return false
}

// Generate a human-friendly version of the board to display in a terminal
func (b Board) Display() string {
	// Use a buffer to efficiently concatenate the board display string
	var display bytes.Buffer
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			display.WriteString(fmt.Sprintf(" %s ", displayMap[b.Tiles[row][col]]))

			// Add vertical separators for internal cells
			if col < 2 {
				display.WriteString("|")
			}
		}
		// Add horizontal separators for internal cells
		if row < 2 {
			display.WriteString("\n-----------\n")
		}
	}
	// Finish with a new line, ready for the next display
	display.WriteString("\n")

	return display.String()
}
