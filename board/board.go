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

	moveMap = map[string]int{
		"X": P1Val,
		"O": P2Val,
	}
)

type Board struct {
	Cells [Size][Size]int
}

// Create a 3x3 two-dimensional array and fill it with the zero value
func (b *Board) Initialise() {
	b.Cells = [Size][Size]int{}
}

func (b *Board) Clear() {
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			b.Cells[row][col] = 0
		}
	}
}

// Determine whether a proposed move is valid by checking for collisions
// with non-zero values
func (b *Board) IsValidMove(x, y int) bool {
	if b.Cells[x][y] != ZeroVal {
		return false
	}

	return true
}

// RegisterMove updates the board's cells with a player's move
func (b *Board) RegisterMove(player string, x, y int) {
	b.Cells[x][y] = moveMap[player]
}

// Determine whether the board has a winner by checking rows, colums and diagonals
func (b *Board) HasWinner() bool {
	// Check rows and cols for non-zero matches
	for i := 0; i < Size; i++ {
		if b.Cells[i][0] != ZeroVal && b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][1] == b.Cells[i][2] {
			return true
		}

		if b.Cells[0][i] != ZeroVal && b.Cells[0][i] == b.Cells[1][i] && b.Cells[1][i] == b.Cells[2][i] {
			return true
		}
	}

	// Check diagonals for non-zero matches
	if b.Cells[1][1] != ZeroVal && b.Cells[0][0] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][2] {
		return true
	}
	if b.Cells[1][1] != ZeroVal && b.Cells[0][2] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][0] {
		return true
	}

	return false
}

// Generate a human-friendly version of the board to display in a terminal
func (b *Board) Display() string {
	// Use a buffer to efficiently concatenate the board display string
	var display bytes.Buffer
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			display.WriteString(fmt.Sprintf(" %s ", displayMap[b.Cells[row][col]]))

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
