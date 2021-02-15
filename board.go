// board.go for CSI 380 Assignment 3
// This file contains types and an interface for maintaining the state of a game.
// They are generic enough to work for several games, but specifically will work for Connect 4.
// You can implement this Assignment without making any changes to this file.

package main

import "fmt"

// Represents the key to a transition from one position
// to another position
// In Connect 4 this is which column is selected
// to drop a piece
type Move uint

// Represents a player's piece and also turns.
type Piece uint

const Black Piece = 1
const Red Piece = 2

func (piece Piece) opposite() Piece {
	if piece == 0 {
		return piece
	}
	return 3 - piece
}

// Description of a piece; useful to be used in the
// description of a board
func (piece Piece) String() string {
	switch piece {
	case Black:
		return "+"
	case Red:
		return "*"
	default:
		return " "
	}
}

// A generic interface that could represent a board (read position)
// in many different board games that you will implicitly
// need to implement in your connect 4 game in the struct C4Board
// Minimax depends on this interface
type Board interface {
	IsWin() bool
	IsDraw() bool
	Evaluate(player Piece) float32
	LegalMoves() []Move
	MakeMove(move Move) Board
	Turn() Piece
	fmt.Stringer
}

// Return if move is in the given list of Moves or not
func contains(list []Move, move Move) bool {
	for _, m := range list {
		if m == move {
			return true
		}
	}
	return false
}
