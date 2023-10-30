package board

const (
	Size = 3

	ZeroVal = 0
	P1Val   = 1
	P2Val   = -1
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

// ...
func (b Board) Display() string {
	return ""
}
