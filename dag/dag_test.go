package dag

import (
	"fmt"
	"testing"
)

func TestDAG_FromEdgeList(t *testing.T) {
	testCases := []struct {
		desc     string
		edgeList [][]int
	}{
		{
			desc: "",
			edgeList: [][]int{
				{1, 2},
			},
		},
		{
			desc: "",
			edgeList: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{1, 5},
				{1, 6},
				{1, 7},
			},
		},
		{
			desc: "",
			edgeList: [][]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{2, 5},
				{3, 6},
				{3, 7},
			},
		},
		{
			desc: "",
			edgeList: [][]int{
				{2, 5},
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 6},
				{3, 7},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dag := FromEdgeList(tC.edgeList)

			dag.BFS(func(n *Node[int]) {
				fmt.Printf("%+v\n", n)
			})
		})
	}
}
