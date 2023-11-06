package board_test

import (
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
	b.InitialiseBoard()
}

func TestBoard(t *testing.T) {
	Convey("Given a board in some state", t, withCleanup(func() {
		type testCase struct {
			description string
		}
		cases := []testCase{
			{
				description: "test case 1",
			},
			{
				description: "test case 2",
			},
		}

		for _, tc := range cases {
			Convey(tc.description, func() {
				So(true, ShouldBeTrue)
			})
		}
	}))
}
