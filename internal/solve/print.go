package solve

import "fmt"

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
	fmt.Printf("Moves: \n")
	for i := range nodes {
		if i == 0 {
			fmt.Printf("Initial state: \n")
			for j := range nodes[len(nodes)-i-1].puzzle {
				fmt.Printf("%v\n", nodes[len(nodes)-i-1].puzzle[j])
			}
		} else {
			fmt.Printf("------MOVE %v ------\n", i+1)
			for j := range nodes[len(nodes)-i-1].puzzle {
				fmt.Printf("%v", nodes[len(nodes)-i-1].puzzle[j])
				fmt.Printf("------>")
				fmt.Printf("%v\n", nodes[len(nodes)-i].puzzle[j])
			}
		}
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
