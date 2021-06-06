package solve

import (
	"fmt"

	"github.com/fatih/color"
)

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

func nodeIsWorth(closedMap map[string]*Node, openMap map[string]*Node, node *Node) bool {
	if found, exists := closedMap[node.puzzleString]; exists {
		return node.cost < found.cost
	}

	if found, exists := openMap[node.puzzleString]; exists {
		return found.cost < node.cost
	}

	return true
}

func ResolvingPath(node *Node) []*Node {
	currentNode := node
	nodes := make([]*Node, 0)

	for currentNode.parent != nil {
		nodes = append(nodes, currentNode)
		currentNode = currentNode.parent
	}

	return nodes
}

func Astar(puzzle [][]int) {
	puzzleSize := len(puzzle)
	zeroPosition := ZeroPosition(puzzle)

	solvedPuzzle, solvedPiecePositions := SolvedPuzzle(puzzleSize)
	if !Solvable(puzzle, solvedPuzzle) {
		color.Set(color.FgRed)
		fmt.Printf("Puzzle is not solvable\n")
		color.Unset()
		return
	}

	rootNode := &Node{
		puzzle:       puzzle,
		puzzleString: PuzzleToString(puzzle),
		cost:         0,
		heuristic:    Heuristic(puzzle, solvedPiecePositions),
		score:        Heuristic(puzzle, solvedPiecePositions),
		zeroPosition: zeroPosition,
		parent:       nil,
	}

	//Stats variables
	selectedStatesCounter := 0
	maximumOpenStates := 0

	//openList hashmap
	openMap := make(map[string]*Node)
	closedMap := make(map[string]*Node)

	openList := Queue{nil, 0}
	openMap[rootNode.puzzleString] = rootNode
	openList.Add(rootNode)
	var node *Node

	for {
		node = openList.Pop()
		delete(openMap,node.puzzleString)
		selectedStatesCounter++

		if openList.size > maximumOpenStates {
			maximumOpenStates = openList.size
		}

		if node == nil {
			break
		}

		if isSame(node.puzzle, solvedPuzzle) {
			path := ResolvingPath(node)
			color.Set(color.FgGreen)
			PrettyResolvingPath(path)
			color.Unset()
			color.Set(color.FgRed)
			fmt.Printf("\nStatistics:\n")
			fmt.Printf("Number of moves:\t\t %v\n", len(path))
			fmt.Printf("Number of states browsed:\t %v\n", selectedStatesCounter)
			fmt.Printf("Maximum number of open states:\t %v\n", maximumOpenStates)
			color.Unset()
			return
		}

		for _, neighbor := range Neighbors(node) {
			neighbor.heuristic = Heuristic(node.puzzle, solvedPiecePositions)
			neighbor.score = neighbor.cost + neighbor.heuristic

			if nodeIsWorth(closedMap, openMap, neighbor) {
				openList.Add(neighbor)
				openMap[neighbor.puzzleString] = neighbor
			}
		}

		closedMap[node.puzzleString] = node
	}

	fmt.Printf("Cannot solve puzzle")
}
