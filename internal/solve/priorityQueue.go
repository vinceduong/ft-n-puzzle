package solve


type QueueElement struct {
	node *Node
	next *QueueElement
}

type Queue struct {
	head *QueueElement
	size int
}

func (q *Queue) Add(node *Node) {
	q.size++

	if q.head == nil {
		q.head = &QueueElement{
			node: node,
			next: nil,
		}
		return
	}

	if (node.score < q.head.node.score) {
		q.head = &QueueElement{
			node: node,
			next: q.head,
		}
		return
	}

	current := q.head

	for current.next != nil {
		if (node.score < current.next.node.score) {
			newElement := &QueueElement{
				node: node,
				next: current.next,
			}

			current.next = newElement

			return

		}
		current = current.next
	}

	current.next = &QueueElement {
		node,
		nil,
	}
}

func (q *Queue) Pop() *Node {
	if (q.head == nil) {
		return nil
	}
	q.size--

	node := q.head.node

	q.head = q.head.next

	return node
}

func (q *Queue) Contains(puzzle [][]int) *Node {
	current := q.head

	for current != nil {
		if isSame(current.node.puzzle, puzzle) {
			return current.node
		}

		current = current.next
	}

	return nil
}

