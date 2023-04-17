package main

import (
	"github.com/leetcode-golang-classroom/timber/timber"
	timber_example "github.com/leetcode-golang-classroom/timber/timber/example"
	timber_interface "github.com/leetcode-golang-classroom/timber/timber/interface"
)

func main() {
	tree := timber_example.Branch{
		Label: "branch",
		Items: []timber_interface.NodeTree{
			timber_example.Leaf{
				Label: "leaf",
			},
			timber_example.Branch{
				Label: "branch",
				Items: []timber_interface.NodeTree{
					timber_example.Leaf{
						Label: "leaf",
					},
					timber_example.Leaf{
						Label: "leaf",
					},
					timber_example.Leaf{
						Label: "leaf",
					},
				},
			},
			timber_example.Leaf{
				Label: "leaf",
			},
		},
	}
	printer := timber.NewDefaultPrinter()
	printer.Print(tree)
}

type Leaf struct {
	label string
}

func (l Leaf) Display() string {
	return l.label
}

func (l Leaf) Components() []timber_interface.NodeTree {
	return nil
}

type Branch struct {
	label      string
	components []timber_interface.NodeTree
}

func (b Branch) Display() string {
	return b.label
}

func (b Branch) Components() []timber_interface.NodeTree {
	return b.components
}
