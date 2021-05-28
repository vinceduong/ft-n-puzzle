package solve

import(
	"fmt"
	"strings"
)

func PrettyPuzzle(puzzle [][]int) {
	fmt.Printf("Puzzle: \n")

	for i := range puzzle {
		fmt.Printf("%v\n", puzzle[i])
	}
}

func PrettyNode(node *Node) {
	fmt.Printf("Node puzzle: \n")

	for i := range node.puzzle {
		fmt.Printf("%v\n", node.puzzle[i])
	}

	fmt.Printf("Node cost: %v\n", node.cost)
	fmt.Printf("Node heuristic: %v\n", node.heuristic)
	fmt.Printf("Node score: %v\n", node.score)
	fmt.Printf("Node parent: %v\n", node.parent)
	fmt.Printf("-------------------------------------\n")
}

func PrettyNodes(nodes []*Node) {
	for _, node := range nodes {
		PrettyNode(node)
	}
}

func PrettyResolvingPath(nodes []*Node) {
	puzzleSize := len(nodes[0].puzzle)
	maxPiece := puzzleSize * puzzleSize - 1

	maxPiecePadding := 0
	for maxPiece > 0 {
		maxPiece /= 10
		maxPiecePadding++
	}

	moves := len(nodes)
	moveNumberPadding := 0
	for moves > 0 {
		moves /= 10
		moveNumberPadding++
	}

	const arrow = "-----> "
	const move = "MOVE"
	const space = "  "

	lineSize := 2 * puzzleSize * (maxPiecePadding + 1) + len(arrow) - 1
	headerSpaceSize := lineSize - len(space) - len(move) - len(space) -  moveNumberPadding
	headerSize := headerSpaceSize / 2
	moveNumberPadding += headerSpaceSize % 2

	header := strings.Repeat("-", headerSize)
	footer := strings.Repeat("-", lineSize)

	for i := range nodes {
		if i == 0 {
			continue
		} else {
			fmt.Printf("%v" + "%v" + "%v" + "%*d" + "%v" + "%v\n",
				header,
				space,
				move,
				moveNumberPadding,
				i,
				space,
				header,
			)
			for j := range nodes[len(nodes)-i-1].puzzle {
				for k := range nodes[len(nodes)-i-1].puzzle[j] {
					fmt.Printf("%*d ", maxPiecePadding,nodes[len(nodes)-i].puzzle[j][k])
				}
				fmt.Printf(arrow)
				for k := range nodes[len(nodes)-i-1].puzzle[j] {
					fmt.Printf("%*d ", maxPiecePadding,nodes[len(nodes)-i-1].puzzle[j][k])
				}
				fmt.Printf("\n")
			}
		}
		fmt.Printf(footer + "\n")
	}
}

func PrettyQueue(q Queue) {
	current := q.head

	fmt.Printf("---- Printing Queue ----\n")
	fmt.Printf("Queue info = %v\n", q)
	i := 0
	for current != nil {
		fmt.Printf("Node number %v\n", i)
		PrettyNode(current.node)
		current = current.next
		i++
	}

	fmt.Printf("---- End of Queue ----\n")
}
