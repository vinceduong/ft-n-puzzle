package main

import (
	"fmt"
	"log"
	"os"
	"sort"

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

func isMissPlaced(row int, column int, pos Position) int {
	if row != pos.row || column != pos.column {
		return 1
	}
	return 0
}

func Heuristic(t string, puzzle [][]int, solvedPiecePositions map[int]Position) int {
	heuristic := 0

	var function func(row int, column int, pos Position) int

	switch t {
	case "manhattan":
		function = ManhattanDistance
	case "mp":
		function = isMissPlaced
	default:
		function = ManhattanDistance
	}

	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			distance := function(i, j, solvedPiecePositions[pieceNumber])
			// distance := isMissPlaced(i, j, solvedPiecePositions[pieceNumber])
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

func prettyNode(node *Node) {
	fmt.Printf("Node puzzle: \n")

	for i := range node.puzzle {
		fmt.Printf("%v\n", node.puzzle[i])
	}

	fmt.Printf("Node cost: %v\n", node.cost)
	fmt.Printf("Node heuristic: %v\n", node.heuristic)
	fmt.Printf("Node parent: %v\n", node.parent)
	fmt.Printf("-------------------------------------\n")
}

func prettyNodes(nodes []*Node) {
	for _, node := range nodes {
		prettyNode(node)
	}
}

func Neighbors(node *Node, solvedPiecePositions map[int]Position) []*Node {
	puzzleSize := len(node.puzzle)
	potentialZeroPositions := PotentialZeroPositions(node.zeroPosition, puzzleSize)
	neighbors := make([]*Node, len(potentialZeroPositions))

	for i, position := range potentialZeroPositions {
		newPuzzle := SwapPuzzlePieces(node.puzzle, node.zeroPosition, position)
		cost := node.cost + 1

		neighbors[i] = &Node{
			newPuzzle,
			cost, 0, 0,
			position,
			node,
		}
	}

	return neighbors
}

func isSame(puzzle1 [][]int, puzzle2 [][]int) bool {
	for i := range puzzle1 {
		for j := range puzzle1[i] {
			if puzzle1[i][j] != puzzle2[i][j] {
				return false
			}
		}
	}

	return true
}

func nodeIsWorth(closedList []*Node, openList []*Node, node *Node) bool {
	// fmt.Println("----------CHILD NODE---------")
	// prettyNode(node)
	for _, n := range closedList {
		if isSame(n.puzzle, node.puzzle) {
			// fmt.Print("Node is not worth\n")
			return false
		}
	}

	for _, n := range openList {
		if isSame(n.puzzle, node.puzzle) && n.cost <= node.cost {
			// fmt.Print("Node is not worth\n")
			return false
		}
	}

	// fmt.Print("Node is worth\n")
	return true
}

func showResolvingPath(node *Node) {
	currentNode := node
	nodes := make([]Node, 0)

	for currentNode.parent != nil {
		nodes = append(nodes, *currentNode)
		currentNode = currentNode.parent
	}

	for i := range nodes {
		fmt.Printf("------STEP %v ------\n\n", i+1)
		for j := range nodes[len(nodes)-i-1].puzzle {
			fmt.Printf("%v\n", nodes[len(nodes)-i-1].puzzle[j])
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No file provided")
	}
	fileName := os.Args[1]

	lines := parse.LinesFromFile(fileName)
	puzzle, puzzleSize := parse.PuzzleFromLines(lines)
	zeroPosition := ZeroPosition(puzzle)

	solvedPuzzle, piecePositions := SolvedPuzzle(puzzleSize)

	rootNode := &Node{
		puzzle,
		0, Heuristic("manhattan", puzzle, piecePositions), Heuristic("manhattan", puzzle, piecePositions),
		zeroPosition,
		nil,
	}

	closedList := make([]*Node, 0)
	openList := make([]*Node, 1)
	openList[0] = rootNode
	var node *Node

	for len(openList) > 0 {
		node, openList = openList[0], openList[1:]

		if isSame(node.puzzle, solvedPuzzle) {
			showResolvingPath(node)
			fmt.Println("Puzzle is solved")
			return
		}
		for _, neighbor := range Neighbors(node, piecePositions) {
			neighbor.heuristic = Heuristic("manhattan", node.puzzle, piecePositions)
			neighbor.score = neighbor.cost + neighbor.heuristic

			if nodeIsWorth(closedList, openList, neighbor) {
				openList = append(openList, neighbor)
			}
		}

		sort.Slice(openList, func(i, j int) bool {
			if openList[i].heuristic == openList[j].heuristic {
				return openList[i].score < openList[j].score
			} else {
				return openList[i].heuristic < openList[j].heuristic
			}
		})

		closedList = append(closedList, node)
	}

	fmt.Printf("Cannot solve puzzle")
}
