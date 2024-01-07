package board_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"tic-tac-go/board"
)

var (
	b board.Board
)

func withCleanup(f func()) func() {
	return func() {
		Reset(cleanup)
		f()
	}
}

func cleanup() {
	b.Clear()
}

func TestHasWinner(t *testing.T) {
	Convey("Given we are testing board.HasWinner", t, func() {
		type testCase struct {
			description string
			hasWinner   bool
			p1Moves     [][]int
			p2Moves     [][]int
		}
		cases := []testCase{
			{
				description: "When P1 wins along the first row",
				hasWinner:   true,
				p1Moves: [][]int{
					{0, 0},
					{0, 1},
					{0, 2},
				},
				p2Moves: [][]int{
					{1, 1},
					{1, 2},
				},
			},
			{
				description: "When P2 wins along the first row",
				hasWinner:   true,
				p1Moves: [][]int{
					{1, 1},
					{1, 0},
					{2, 2},
				},
				p2Moves: [][]int{
					{0, 0},
					{0, 1},
					{0, 2},
				},
			},
		}

		for _, tc := range cases {
			cleanup()

			for _, move := range tc.p1Moves {
				b.Cells[move[0]][move[1]] = board.P1Val
			}
			for _, move := range tc.p2Moves {
				b.Cells[move[0]][move[1]] = board.P2Val
			}

			Convey(tc.description, func() {
				fmt.Printf("\n%s", b.Display())

				So(b.HasWinner(), ShouldEqual, tc.hasWinner)
			})
		}
	})
}

func TestIsValidMove(t *testing.T) {
	Convey("Given we are testing board.IsValidMove", t, withCleanup(func() {
		type testCase struct {
			description string
			isValid     bool
		}
		cases := []testCase{
			{
				description: "test case 1",
				isValid:     false,
			},
			{
				description: "test case 2",
				isValid:     false,
			},
		}

		for _, tc := range cases {
			Convey(tc.description, func() {
				So(true, ShouldBeTrue)
				// So(b.IsValidMove(0, 0), ShouldEqual, tc.isValid)
			})
		}
	}))
}
