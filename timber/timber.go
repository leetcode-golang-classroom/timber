package timber

import (
	timber_interface "github.com/leetcode-golang-classroom/timber/timber/interface"
	log "github.com/sirupsen/logrus"
)

func Print(tree timber_interface.NodeTree) {
	doPrint(tree, node{prefix: "", isLast: true})
}

func doPrint(tree timber_interface.NodeTree, node node) {
	log.Print(node.display(), tree.Display())
	length := len(tree.Components())
	for idx, child := range tree.Components() {
		isLast := idx == length-1
		doPrint(child, node.createChildNode(isLast))
	}
}

// node
type node struct {
	prefix string
	isLast bool
}

func (n node) display() string {
	if n.isLast {
		return n.prefix + "\u2514 "
	}
	return n.prefix + "\u251c "
}

func (n node) createChildNode(isLast bool) node {
	return node{
		prefix: n.prefix + n.nextPrefix(),
		isLast: isLast,
	}
}

func (n node) nextPrefix() string {
	if n.isLast {
		return " "
	}
	return "\u2502 "
}
