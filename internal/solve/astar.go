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

	if similarNode != nil && similarNode.cost < node.cost  {
		return false
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
	openList := Queue{nil, 0}
	openList.Add(rootNode)
	var node *Node

	for /*i := 0; i < 10; i++*/ {

//		PrettyQueue(openList)
		node = openList.Pop()
//		fmt.Printf("openList = %v\n", openList)

		if node == nil {
			break
		}
		if isSame(node.puzzle, solvedPuzzle) {
			ShowResolvingPath(node)
			fmt.Println("Puzzle is solved")
			return
		}
		for _, neighbor := range Neighbors(node, solvedPiecePositions) {
			neighbor.heuristic = Heuristic("manhattan", node.puzzle, solvedPiecePositions)
			neighbor.score = neighbor.cost + neighbor.heuristic
//			fmt.Printf("neighbor = %v\n", neighbor);
//			fmt.Printf("neighbor heuristic: %v\n", neighbor.heuristic);
//			fmt.Printf("neighbor score: %v\n", neighbor.score);

			if nodeIsWorth(closedList, openList, neighbor) {
//				fmt.Printf("Node is worth!!\n")
				openList.Add(neighbor)
			}
		}

		closedList = append(closedList, node)
	}

	fmt.Printf("Cannot solve puzzle")
}
