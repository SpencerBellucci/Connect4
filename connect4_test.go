// Connect 4 for CSI 380 by David Kopec
// This file tests the final implementation
// with test positions as well as a couple of the
// methods of C4Board.

// Note that in board literals, columns go from left to right,
// not up and down.

package main

import "testing"

// Check if two slices of moves are equivalent
func checkEquivalent(p1s, p2s []Move) bool {
	if len(p1s) != len(p2s) { // same length
		return false
	}

	for _, p := range p1s { // same items
		if !contains(p2s, p) {
			return false
		}
	}

	return true
}

// Test if the right legal moves are generated
// blank board
func TestLegalMoves1(t *testing.T) {
	b := C4Board{turn: Black}
	expected := []Move{0, 1, 2, 3, 4, 5, 6}
	actual := b.LegalMoves()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test if the right legal moves are generated
// board with 1 filled column
func TestLegalMoves2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 6, 2, 5, 0, 0},
		turn:     Red}
	expected := []Move{0, 1, 3, 4, 5, 6}
	actual := b.LegalMoves()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test if a non-full board is called a draw
func TestDraw1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 6, 2, 5, 0, 0},
		turn:     Red}
	expected := false
	actual := b.IsDraw()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test if a full board is called a draw
func TestDraw2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 1, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsDraw()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// false
func TestWin1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 1, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := false
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 1, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin3(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 1, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 1},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 1, 2}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin4(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 1},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 2, 1, 1, 1, 2}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin5(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 2, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{4, 3, 1, 0, 0, 0, 0},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test Concurrent Find Best Move
// stop the win
func TestConcurrentFindBestMove1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 0, 0, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{2, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{2, 3, 1, 1, 0, 0, 0},
		turn:     Red}
	expected := Move(1)
	actual := ConcurrentFindBestMove(b, 4)
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test Concurrent Find Best Move
// take the win
func TestConcurrentFindBestMove2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 0, 0, 0, 0},
		[6]Piece{1, 1, 0, 0, 0, 0},
		[6]Piece{1, 1, 0, 0, 0, 0},
		[6]Piece{2, 2, 1, 0, 0, 0},
		[6]Piece{2, 1, 2, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{2, 2, 2, 3, 3, 0, 0},
		turn:     Black}
	expected := Move(4)
	actual := ConcurrentFindBestMove(b, 3)
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test Concurrent Find Best Move
// block a future win
func TestConcurrentFindBestMove3(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 1, 2, 0, 0, 0},
		turn:     Red}
	actual := ConcurrentFindBestMove(b, 5)
	if (actual != Move(1)) && (actual != Move(4)) { // two save moves
		t.Errorf("Test failed: expected %v to be 1 or 4", actual)
	}
}
