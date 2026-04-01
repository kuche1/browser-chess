package lib

type Piece uint8

const (
	PieceKing Piece = iota
	PieceQueen
	PieceRook
	PieceBishop
	PieceKnight
	PiecePawn
)
