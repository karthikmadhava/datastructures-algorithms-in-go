package DataStructures

import "fmt"

// TreeNode represents a node in the tree.
type TreeNode struct {
	Name     string
	Children []*TreeNode
}

// NewTreeNode creates a new TreeNode with the given name.
func NewTreeNode(name string) *TreeNode {
	return &TreeNode{
		Name:     name,
		Children: []*TreeNode{},
	}
}

// AddChild adds a child node to the current node.
func (n *TreeNode) AddChild(child *TreeNode) {
	n.Children = append(n.Children, child)
}

// PrintTree prints the tree structure recursively.
func (n *TreeNode) PrintTree(level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%s- %s\n", indent, n.Name)
	for _, child := range n.Children {
		child.PrintTree(level + 1)
	}
}

// CreateTree creates a sample tree
func CreateTree() *TreeNode {
	//root
	root := NewTreeNode("Root")

	//level 1
	child1 := NewTreeNode("Child1")
	child2 := NewTreeNode("Child2")

	//level 2
	grandchild1 := NewTreeNode("Grandchild1")
	grandchild2 := NewTreeNode("Grandchild2")
	grandchild3 := NewTreeNode("Grandchild3")

	//adding children
	root.AddChild(child1)
	root.AddChild(child2)
	child1.AddChild(grandchild1)
	child1.AddChild(grandchild2)
	child2.AddChild(grandchild3)
	return root

}
