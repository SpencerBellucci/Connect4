// main.go for CSI 380 Assignment 3
// This file includes the main game loop
// that actually creates a human vs computer game.

package main

import "fmt"

var gameBoard Board = C4Board{turn: Black}

func getPlayerMove() Move {
	var playerMove Move = 1000
	for !contains(gameBoard.LegalMoves(), playerMove) {
		println("Enter a column (0-6):")
		var input int
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			println(err)
			continue
		}
		playerMove = Move(input)
	}
	return playerMove
}

func main() {
	// Main game loop
	for {
		// Make human move
		humanMove := getPlayerMove()
		gameBoard = gameBoard.MakeMove(humanMove)

		if gameBoard.IsWin() {
			println("Human wins!")
			break
		} else if gameBoard.IsDraw() {
			println("Draw!")
			break
		}
		computerMove := ConcurrentFindBestMove(gameBoard, 5)
		fmt.Printf("Computer move is %d\n", computerMove)
		gameBoard = gameBoard.MakeMove(computerMove)
		print(gameBoard.String())
		if gameBoard.IsWin() {
			println("Computer wins!")
			break
		} else if gameBoard.IsDraw() {
			println("Draw!")
			break
		}
	}
}

