package solve

import "fmt"

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

func ShowResolvingPath(node *Node) {
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
