package dag

type Node[T any] struct {
	Value    T
	Parent   *Node[T]
	Children []*Node[T]
}

type DAG[T any] struct {
	Head *Node[T]
}

func FromEdgeList[T comparable](edges [][]T) *DAG[T] {
	nodes := make(map[T]*Node[T])

	for _, e := range edges {
		fromVal := e[0]
		from, ok := nodes[fromVal]
		if !ok {
			from = &Node[T]{
				Value:  fromVal,
				Parent: nil,
			}
		}

		toVal := e[1]
		to, ok := nodes[toVal]
		if !ok {
			to = &Node[T]{
				Value:  toVal,
				Parent: nil,
			}
		}

		to.Parent = from
		from.Children = append(from.Children, to)

		nodes[fromVal] = from
		nodes[toVal] = to
	}

	for _, node := range nodes {
		if node.Parent == nil {
			return &DAG[T]{
				Head: node,
			}
		}
	}

	panic("Unreachable: no head node found")
}

func (dag *DAG[T]) BFS(operation func(*Node[T])) {
	var toVisit []*Node[T]

	dag.Head.BFS(operation, toVisit)
}

func (node *Node[T]) BFS(op func(*Node[T]), toVisit []*Node[T]) {
	op(node) // first, do the thing

	// add the children nodes to the queue
	if len(node.Children) > 0 {
		toVisit = append(toVisit, node.Children...)
	}

	// if we're done, return
	if len(toVisit) == 0 {
		return
	}

	// otherwise dequeue and invoke bfs
	next := toVisit[0]
	next.BFS(op, toVisit[1:])
}
