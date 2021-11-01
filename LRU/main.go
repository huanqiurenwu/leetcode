package LRU

type Node struct {
	key, val int
	next, prev *Node
}

type NodeList struct {
	removeLast
}