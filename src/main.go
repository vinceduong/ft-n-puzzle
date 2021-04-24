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
	zeroPosition           Position
	parent                 *Node
}

type Position struct {
	row    int
	column int
}

func getSolvedPuzzle(puzzleSize int) ([][]int, map[int]Position) {
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

	piecePositions := make(map[int]Position)

	for pieceNumber <= numberOfPieces {
		if pieceNumber != numberOfPieces {
			puzzle[row][column] = pieceNumber
			piecePositions[pieceNumber] = Position{row, column}
			pieceNumber++
		} else {
			puzzle[row][column] = 0
			piecePositions[0] = Position{row, column}
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

	return puzzle, piecePositions
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(row int, column int, pos Position) int {
	return Abs(row-pos.row) + Abs(column-pos.column)
}

func heuristic(puzzle [][]int, solvedPiecePositions map[int]Position) int {
	heuristic := 0

	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			distance := manhattanDistance(i, j, solvedPiecePositions[pieceNumber])
			heuristic += distance
		}
	}

	return heuristic
}

func potentialZeroPositions(zeroPosition Position, puzzleSize int) []Position {
	var positions []Position

	if zeroPosition.row > 0 {
		positions = append(
			positions,
			Position{zeroPosition.row - 1, zeroPosition.column},
		)
	}

	if zeroPosition.row < puzzleSize-1 {
		positions = append(
			positions,
			Position{zeroPosition.row + 1, zeroPosition.column},
		)
	}

	if zeroPosition.column > 0 {
		positions = append(
			positions,
			Position{zeroPosition.row, zeroPosition.column - 1},
		)
	}

	if zeroPosition.column < puzzleSize-1 {
		positions = append(
			positions,
			Position{zeroPosition.row, zeroPosition.column + 1},
		)
	}

	return positions
}

func getZeroPosition(puzzle [][]int) Position {
	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			if pieceNumber == 0 {
				return Position{i, j}
			}
		}
	}

	return Position{}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No file provided")
	}
	fileName := os.Args[1]

	lines := parse.GetLinesFromFile(fileName)
	fmt.Println("'" + strings.Join(lines, `','`) + `'`)
	puzzle, puzzleSize := parse.GetPuzzleFromLines(lines)
	zeroPosition := getZeroPosition(puzzle)

	fmt.Printf("Puzzle: %v\n\n", puzzle)
	fmt.Printf("Zero position: %v\n\n", zeroPosition)

	solvedPuzzle, piecePositions := getSolvedPuzzle(puzzleSize)
	fmt.Printf("Solved Puzzle: %v\n\n", solvedPuzzle)
	fmt.Printf("Heuristic: %v\n\n", heuristic(solvedPuzzle, piecePositions))
	fmt.Printf("Neighboors: %v\n", potentialZeroPositions(zeroPosition, puzzleSize))
}
