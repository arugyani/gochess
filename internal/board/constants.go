package board

type Color int

const (
	White Color = iota
	Black
	Colors
)

var colorName = map[Color]string{
	White: "white",
	Black: "black",
}

func (c Color) String() string {
	return colorName[c]
}

type Piece int

const (
	Pawn Piece = iota
	Rook
	Knight
	Bishop
	Queen
	King
	Pieces
)

var pieceName = map[Piece]string{
	Pawn:   "pawn",
	Rook:   "rook",
	Knight: "knight",
	Bishop: "bishop",
	Queen:  "queen",
	King:   "king",
}

func (p Piece) String() string {
	return pieceName[p]
}

const (
	NumSquares        = 64
	WhitePawnsStart   = 0x000000000000FF00
	WhiteRooksStart   = 0x0000000000000081
	WhiteKnightsStart = 0x0000000000000042
	WhiteBishopsStart = 0x0000000000000024
	WhiteQueenStart   = 0x0000000000000010
	WhiteKingStart    = 0x0000000000000008
	BlackPawnsStart   = 0x00FF000000000000
	BlackRooksStart   = 0x8100000000000000
	BlackKnightsStart = 0x4200000000000000
	BlackBishopsStart = 0x2400000000000000
	BlackQueenStart   = 0x1000000000000000
	BlackKingStart    = 0x0800000000000000
)
