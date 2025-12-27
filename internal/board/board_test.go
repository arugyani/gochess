package board

import "testing"

func TestBoardExists(t *testing.T) {
	got := NewBoard()

	if got == nil {
		t.Fatal("Expected board to be initialized, but got nil")
	}
}

func TestGetPiece(t *testing.T) {
	board := NewBoard()

	t.Run("In Bounds", func(t *testing.T) {
		got := board.GetPiece(55)
		want := Square{
			Piece:    Pawn,
			Color:    Black,
			Occupied: true,
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
