package solve

type Node struct {
	puzzle                 [][]int
	puzzleString					 string
	cost, heuristic, score int
	zeroPosition           Position
	parent                 *Node
}

type Position struct {
	row    int
	column int
}
