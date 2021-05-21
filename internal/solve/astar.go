package solve

import (
	"fmt"
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

func nodeIsWorth(closedList []*Node, openList Queue, node *Node) bool {
	for _, n := range closedList {
		if isSame(n.puzzle, node.puzzle) {
			return false
		}
	}

	similarNode := openList.Contains(node.puzzle)

	if similarNode != nil && similarNode.cost < node.cost {
		return false
	}

	return true
}

func Astar(puzzle [][]int) {
	puzzleSize := len(puzzle)
	zeroPosition := ZeroPosition(puzzle)

	solvedPuzzle, solvedPiecePositions := SolvedPuzzle(puzzleSize)

	rootNode := &Node{
		puzzle:       puzzle,
		cost:         0,
		heuristic:    Heuristic("manhattan", puzzle, solvedPiecePositions),
		score:        Heuristic("manhattan", puzzle, solvedPiecePositions),
		zeroPosition: zeroPosition,
		parent:       nil,
	}

	//openMap := make(map[string]*Node)
	//closedMap := make(map[string]*Node)

	closedList := make([]*Node, 0)
	openList := Queue{nil, 0}
	openList.Add(rootNode)
	var node *Node

	for {

		node = openList.Pop()

		if node == nil {
			break
		}
		if isSame(node.puzzle, solvedPuzzle) {
			ShowResolvingPath(node)
			fmt.Println("Puzzle is solved")
			return
		}
		for _, neighbor := range Neighbors(node) {
			neighbor.heuristic = Heuristic("manhattan", node.puzzle, solvedPiecePositions)
			neighbor.score = neighbor.cost + neighbor.heuristic

			if nodeIsWorth(closedList, openList, neighbor) {
				openList.Add(neighbor)
			}
		}

		closedList = append(closedList, node)
	}

	fmt.Printf("Cannot solve puzzle")
}
