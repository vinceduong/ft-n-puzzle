package solve

import (
	"fmt"
	"sort"
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

func nodeIsWorth(closedList []*Node, openList []*Node, node *Node) bool {
	for _, n := range closedList {
		if isSame(n.puzzle, node.puzzle) {
			return false
		}
	}

	for _, n := range openList {
		if isSame(n.puzzle, node.puzzle) && n.cost <= node.cost {
			return false
		}
	}

	return true
}

func Astar(puzzle [][]int) {
	puzzleSize := len(puzzle)
	zeroPosition := ZeroPosition(puzzle)

	solvedPuzzle, solvedPiecePositions := SolvedPuzzle(puzzleSize)

	rootNode := &Node{
		puzzle,
		0, Heuristic("manhattan", puzzle, solvedPiecePositions),
		Heuristic("manhattan", puzzle, solvedPiecePositions),
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
			ShowResolvingPath(node)
			fmt.Println("Puzzle is solved")
			return
		}
		for _, neighbor := range Neighbors(node, solvedPiecePositions) {
			neighbor.heuristic = Heuristic("manhattan", node.puzzle, solvedPiecePositions)
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
