package main

type Node struct {
	Index     int32
	Id        int64
	NextNodes []*Node
}

func createNodes(graphNodes int32, graphFrom []int32, graphTo []int32, ids []int64) []Node {

}

// return next nodes' indice.
func getNextNodes(graphFrom []int32, graphTo []int32, ids []int64, val int32) []int32 {

}

// return true if there is a same color node.
func existId(ids []int64, id int32) bool {

}

func loadNodes(graphFrom []int32, graphTo []int32, ids []int64) []*Node {

}

// find a same color node.
// if a same color node doesn't exist, check the next nodes.
// if there no same color node after checking all node, return false.
func CheckSameColorNode(nodes []*Node, val int32, distance int32) bool {
	// check whether one of adjacent nodes has same color.

	// if there's a same color node, return true.

	// if there's no same color node, return false.
	return false
}

// this is expected to be called resursively.
// check nearest nodes. and check one more away from checked ones again and again until find the clone.
func findShortestInner(node Node, val int32) int32 {
	//

	// if a same color node is not found,

}

// Complete the findShortest function below.

/*
 * For the unweighted graph, <name>:
 *
 * 1. The number of nodes is <name>Nodes.
 * 2. The number of edges is <name>Edges.
 * 3. An edge exists between <name>From[i] to <name>To[i].
 *
 */
func findShortest(graphNodes int32, graphFrom []int32, graphTo []int32, ids []int64, val int32) int32 {
	// create a Node slice.

	// extract same color nodes.

	// for each node...

	// start to find.

	return 0
}
