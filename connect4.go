// connect4.go for CSI 380 Assignment 3
// The struct C4Board should implement the Board
// interface specified in Board.go
// Note: you will almost certainly need to add additional
// utility functions/methods to this file.

package main

import (
	"fmt"
	// "strings"
)

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
	// create copy of old board
	newBoard := board

	// create new piece, and add to board
	newPiece := newBoard.turn
	newBoard.position[col][(numRows - 1) - newBoard.colCount[col]] = newPiece

	// update board info
	if newBoard.colCount[col] != 5 {
		newBoard.colCount[col] = newBoard.colCount[col] + 1
	}
	newBoard.turn = newBoard.turn.opposite()

	// return the new board
	return newBoard
}

// All of the current legal moves.
// Remember, a move is just the column you can play.
func (board C4Board) LegalMoves() []Move {
	// slice to store possible moves
	var moves []Move

	// check for columns that aren't full
	for i := 0; i < int(numCols); i++ {
		if board.colCount[i] < numCols {
			// add column to legal moves
			moves = append(moves, Move(i))
		}
	}
	return moves
}

// Is it a win?
func (board C4Board) IsWin() bool {
	// see how many in a row
	winBlack := 0
	winRed := 0

	// iterate through slice
	for i := 0; i < len(allSegments); i++ {
		segment := allSegments[i]
		// check if black win
		for j := 0; j < len(segment); j++ {
			// iterate through positions in individual segments
			currentPiece := board.position[segment[j].col][segment[j].row]
			if currentPiece == Black {
				winBlack++
			}
		}
		// check if black has 4 in a row
		if winBlack == 4 {
			// it's a win
			return true
		}
		// if not, set black back to 0
		winBlack = 0

		// check if red win
		for j := 0; j < len(segment); j++ {
			// iterate through positions in individual segments
			currentPiece := board.position[segment[j].col][segment[j].row]
			if currentPiece == Red {
				winRed++
			}
		}
		// check if red has 4 in a row
		if winRed == 4 {
			// it's a win
			return true
		}
		// if not, set black back to 0
		winRed = 0
	}
	// not a win
	return false
}

// Is it a draw?
func (board C4Board) IsDraw() bool {
	// if it's not a win, and no more moves
	if !board.IsWin() && len(board.LegalMoves()) == 0 {
		// it's a draw
		return true
	}
	// not a draw
	return false
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
	var playerScore float32 = 0
	var playerColor Piece

	// determine player color
	if player == Black {
		playerColor = Black
	} else {
		playerColor = Red
	}

	// iterate through segments awarding points
	for i := 0; i < len(allSegments); i++ {
		segment := allSegments[i]

		sameColor := 0
		// check if all red
		for j := 0; j < len(segment); j++ {
			currentPiece := board.position[segment[j].col][segment[j].row]
			// if red, add point to red
			if currentPiece == playerColor{
				sameColor++
			} else if currentPiece == playerColor.opposite() {
				sameColor = 0
				break
			}
		}
		// award points
		if sameColor != 0 {
			switch sameColor {
			case 1:
				playerScore = playerScore + 20
			case 2:
				playerScore = playerScore + 100
			case 3:
				playerScore = playerScore + 500
			case 4:
				playerScore = playerScore + 1000
			}
		}
	}
	return playerScore
}

// Nice to print board representation
// This will be used in play.go to print out the state of the position
// to the user
func (board C4Board) String() string {
	// string to hold the board string
	var drawnBoard string

	// print out row nums for user interface
	fmt.Println("  0   1   2   3   4   5   6")
	// iterate row by row
	for i := 0; i < int(numRows); i++ {
		// iterate through columns
		for j:= 0; j < int(numCols); j++ {
			// print what piece is at the current location
			drawnBoard = drawnBoard + "| " + board.position[j][i].String() + " "
		}
		drawnBoard = drawnBoard + "|\n"
	}

	return drawnBoard
}
