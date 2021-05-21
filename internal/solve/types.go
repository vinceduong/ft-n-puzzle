package solve

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
