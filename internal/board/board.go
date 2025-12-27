package board

import (
	"fmt"
	"gochess/internal/util"
)

type Board struct {
	State [Colors][Pieces]uint64

	AllPieces uint64
}

type Square struct {
	Piece    Piece
	Color    Color
	Occupied bool
}

func (s Square) String() string {
	return fmt.Sprintf("Square{%v, %v, %v}", s.Piece, s.Color, s.Occupied)
}

func (b *Board) GetPiece(square int) Square {
	if square < 0 || square >= NumSquares {
		panic((fmt.Sprintf("square %d out of bounds", square)))
	}

	squareBit := uint64(1 << square)

	isEmpty := (b.AllPieces & squareBit) == 0
	if !isEmpty {
		for color := White; color < Colors; color++ {
			for piece := Pawn; piece < Pieces; piece++ {
				if b.State[color][piece]&squareBit != 0 {
					return Square{Piece: piece, Color: color, Occupied: true}
				}
			}
		}
	}

	return Square{Occupied: false}
}

func NewBoard() *Board {
	b := Board{
		State: [Colors][Pieces]uint64{
			{
				WhitePawnsStart,
				WhiteRooksStart,
				WhiteKnightsStart,
				WhiteBishopsStart,
				WhiteQueenStart,
				WhiteKingStart,
			},
			{
				BlackPawnsStart,
				BlackRooksStart,
				BlackKnightsStart,
				BlackBishopsStart,
				BlackQueenStart,
				BlackKingStart,
			},
		},
	}

	b.AllPieces = util.Collapse(b.State[White][:]) | util.Collapse(b.State[Black][:])

	return &b
}
