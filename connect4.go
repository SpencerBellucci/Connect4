// connect4.go for CSI 380 Assignment 3
// The struct C4Board should implement the Board
// interface specified in Board.go
// Note: you will almost certainly need to add additional
// utility functions/methods to this file.

package main

import "strings"

// size of the board
const numCols uint = 7
const numRows uint = 6

// size of a winning segment in Connect 4
const segmentLength uint = 4

// represents a place on the board
type location struct {
	col uint
	row uint
}

// Represents a segment of 4 locations on the board.
type segment [segmentLength]location

// Returns a slice containing all of the possible segments
// on the board.
func generateSegments() []segment {
	var segments []segment
	// generate vertical segments
	for c := uint(0); c < numCols; c++ {
		for r := uint(0); r < numRows-segmentLength+1; r++ {
			s := segment{{c, r}, {c, r + 1}, {c, r + 2}, {c, r + 3}}
			segments = append(segments, s)
		}
	}
	// generate horizontal segments
	for c := uint(0); c < numCols-segmentLength+1; c++ {
		for r := uint(0); r < numRows; r++ {
			s := segment{{c, r}, {c + 1, r}, {c + 2, r}, {c + 3, r}}
			segments = append(segments, s)
		}
	}
	// generate the bottom left to top right diagonal segments
	for c := uint(0); c < numCols-segmentLength+1; c++ {
		for r := uint(0); r < numRows-segmentLength+1; r++ {
			s := segment{{c, r}, {c + 1, r + 1}, {c + 2, r + 2}, {c + 3, r + 3}}
			segments = append(segments, s)
		}
	}
	// generate the top left to bottom right diagonal segments
	for c := uint(0); c < numCols-segmentLength+1; c++ {
		for r := segmentLength - 1; r < numRows; r++ {
			s := segment{{c, r}, {c + 1, r - 1}, {c + 2, r - 2}, {c + 3, r - 3}}
			segments = append(segments, s)
		}
	}
	return segments
}

var allSegments []segment = generateSegments()

// The main struct that should implement the Board interface
// It maintains the position of a game
// You should not need to add any additional properties to this struct, but
// you may add additional methods
type C4Board struct {
	position [numCols][numRows]Piece  // the grid in Connect 4
	colCount [numCols]uint // how many pieces are in a given column (or how many are "non-empty")
	turn     Piece // who's turn it is to play
}

// Who's turn is it?
func (board C4Board) Turn() Piece {
	return board.turn
}

// Put a piece in column col.
// Returns a copy of the board with the move made.
// Does not check if the column is full (assumes legal move).
func (board C4Board) MakeMove(col Move) Board {
	// YOUR CODE HERE
}

// All of the current legal moves.
// Remember, a move is just the column you can play.
func (board C4Board) LegalMoves() []Move {
	// YOUR CODE HERE
}

// Is it a win?
func (board C4Board) IsWin() bool {
	// YOUR CODE HERE
}

// Is it a draw?
func (board C4Board) IsDraw() bool {
	// YOUR CODE HERE
}

// Who is winning in this position?
// This function scores the position for player
// and returns a numerical score
// When player is doing well, the score should be higher
// When player is doing worse, player's returned score should be lower
// Scores mean nothing except in relation to one another; so you can
// use any scale that makes sense to you
// The more accurately Evaluate() scores a position, the better that minimax will work
// There may be more than one way to evaluate a position but an obvious route
// is to count how many 1 filled, 2 filled, and 3 filled segments of the board
// that the player has (that don't include any of the opponents pieces) and give
// a higher score for 3 filleds than 2 filleds, 1 filleds, etc.
// You may also need to score wins (4 filleds) as very high scores and losses (4 filleds
// for the opponent) as very low scores
// You may want to make helper functions/methods like evaluateSegment() and countSegment(), 
// but it's up to you
func (board C4Board) Evaluate(player Piece) float32 {
	// YOUR CODE HERE
}

// Nice to print board representation
// This will be used in play.go to print out the state of the position
// to the user
func (board C4Board) String() string {
	// YOUR CODE HERE
}
