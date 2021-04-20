package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vinceduong/n-puzzle/src/parse"
)

type Node struct {
	puzzle                 [][]int
	cost, heuristic, score int
}

func getSolvedPuzzle(puzzleSize int) [][]int {
	puzzle := make([][]int, puzzleSize)
	for i := range puzzle {
		puzzle[i] = make([]int, puzzleSize)
		for j := range puzzle[i] {
			puzzle[i][j] = -1
		}
	}

	var (
		numberOfPieces   = puzzleSize * puzzleSize
		pieceNumber      = 1
		row              = 0
		column           = 0
		rowDirection     = 1
		columnDirection  = 1
		incrementingAxis = "column"
	)

	for pieceNumber <= numberOfPieces {
		if pieceNumber != numberOfPieces {
			puzzle[row][column] = pieceNumber
			pieceNumber++
		} else {
			puzzle[row][column] = 0
			pieceNumber++

			continue
		}

		if incrementingAxis == "column" {
			if columnDirection == 1 && (column == puzzleSize-1 || puzzle[row][column+1] != -1) {
				rowDirection = 1
				incrementingAxis = "row"
			}

			if columnDirection == -1 && (column == 0 || puzzle[row][column-1] != -1) {
				rowDirection = -1
				incrementingAxis = "row"
			}
		}

		if incrementingAxis == "row" {
			if rowDirection == 1 && (row == puzzleSize-1 || puzzle[row+1][column] != -1) {
				columnDirection = -1
				incrementingAxis = "column"
			}

			if rowDirection == -1 && (row == 0 || puzzle[row-1][column] != -1) {
				columnDirection = 1
				incrementingAxis = "column"
			}
		}

		if incrementingAxis == "column" {
			column += columnDirection
		} else {
			row += rowDirection
		}
	}

	return puzzle
}

// func getPuzzleHeuristic(puzzle [][]int) int {

// }

// func createNode(puzzle [][]int, cost int) Node {

// }

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No file provided")
	}
	fileName := os.Args[1]

	lines := parse.GetLinesFromFile(fileName)
	fmt.Println("'" + strings.Join(lines, `','`) + `'`)
	puzzle, puzzleSize := parse.GetPuzzleFromLines(lines)
	fmt.Printf("Puzzle: %#v\n", puzzle)
	solvedPuzzle := getSolvedPuzzle(puzzleSize)
	fmt.Printf("Solved Puzzle: %#v\n", solvedPuzzle)
}
