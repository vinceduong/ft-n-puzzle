package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vinceduong/n-puzzle/src/parse"
)

type Node struct {
	puzzle          [][]int
	cost, heuristic int
	zeroPosition    Position
	parent          *Node
}

type Position struct {
	row    int
	column int
}

func SolvedPuzzle(puzzleSize int) ([][]int, map[int]Position) {
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

func ManhattanDistance(row int, column int, pos Position) int {
	return Abs(row-pos.row) + Abs(column-pos.column)
}

func Heuristic(puzzle [][]int, solvedPiecePositions map[int]Position) int {
	heuristic := 0

	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			distance := ManhattanDistance(i, j, solvedPiecePositions[pieceNumber])
			heuristic += distance
		}
	}

	return heuristic
}

func PotentialZeroPositions(zeroPosition Position, puzzleSize int) []Position {
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

func ZeroPosition(puzzle [][]int) Position {
	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			if pieceNumber == 0 {
				return Position{i, j}
			}
		}
	}

	return Position{}
}

func CopyPuzzle(puzzle [][]int) [][]int {
	puzzleCopy := make([][]int, len(puzzle))
	for i := range puzzle {
		puzzleCopy[i] = make([]int, len(puzzle[i]))
		copy(puzzleCopy[i], puzzle[i])
	}

	return puzzleCopy
}

func SwapPuzzlePieces(puzzle [][]int, p1 Position, p2 Position) [][]int {
	newPuzzle := CopyPuzzle(puzzle)

	tmp := newPuzzle[p1.row][p1.column]
	newPuzzle[p1.row][p1.column] = newPuzzle[p2.row][p2.column]
	newPuzzle[p2.row][p2.column] = tmp

	return newPuzzle
}

func prettyNode(node Node) {
	fmt.Printf("Node puzzle: \n")

	for i := range node.puzzle {
		fmt.Printf("%v\n", node.puzzle[i])
	}

	fmt.Printf("Node cost: %v\n", node.cost)
	fmt.Printf("Node heuristic: %v\n", node.heuristic)
	fmt.Printf("Node parent: %v\n", node.parent)
	fmt.Printf("-------------------------------------\n")
}

func prettyNodes(nodes []Node) {
	for _, node := range nodes {
		prettyNode(node)
	}
}

func Neighbors(node *Node, solvedPiecePositions map[int]Position) []Node {
	puzzleSize := len(node.puzzle)
	potentialZeroPositions := PotentialZeroPositions(node.zeroPosition, puzzleSize)
	neighbors := make([]Node, len(potentialZeroPositions))

	for i, position := range potentialZeroPositions {
		newPuzzle := SwapPuzzlePieces(node.puzzle, node.zeroPosition, position)
		cost := node.cost + 1
		heuristic := Heuristic(newPuzzle, solvedPiecePositions) + cost

		neighbors[i] = Node{
			newPuzzle,
			cost, heuristic,
			position,
			node,
		}
	}

	return neighbors
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No file provided")
	}
	fileName := os.Args[1]

	lines := parse.LinesFromFile(fileName)
	fmt.Println("'" + strings.Join(lines, `','`) + `'`)
	puzzle, puzzleSize := parse.PuzzleFromLines(lines)
	zeroPosition := ZeroPosition(puzzle)

	fmt.Printf("Puzzle: %v\n\n", puzzle)
	fmt.Printf("Zero position: %v\n\n", zeroPosition)

	solvedPuzzle, piecePositions := SolvedPuzzle(puzzleSize)
	fmt.Printf("Solved Puzzle: %v\n\n", solvedPuzzle)
	fmt.Printf("Heuristic: %v\n\n", Heuristic(solvedPuzzle, piecePositions))
	fmt.Printf("New Puzzle: %v\n", PotentialZeroPositions(zeroPosition, puzzleSize))
	fmt.Printf("Old Puzzle: %v\n", puzzle)
	fmt.Printf("New Puzzle: %v\n", SwapPuzzlePieces(puzzle, zeroPosition, PotentialZeroPositions(zeroPosition, puzzleSize)[0]))
	fmt.Printf("Old Puzzle: %v\n", puzzle)

	rootNode := Node{
		puzzle,
		0, Heuristic(puzzle, piecePositions),
		zeroPosition,
		nil,
	}

	prettyNode(rootNode)

	prettyNodes(Neighbors(&rootNode, piecePositions))
}
