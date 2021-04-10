package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func tree(node *Node, value int) {
	if node == nil {
		return
	}
	if node.Value < value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
			fmt.Println("left=>", value)
		} else {
			tree(node.Left, value)
		}
	} else {
		if node.Right == nil {
			fmt.Println("right=>", value)
			node.Right = &Node{Value: value}
		} else {
			tree(node.Right, value)
		}
	}
}

func rethree(node *Node) {
	if node == nil {
		return
	}
	if node.Left != nil || node.Right != nil {
		fmt.Println("retree=>", node.Value)

		node.Left, node.Right = node.Right, node.Left
		rethree(node.Left)
		rethree(node.Right)
	}
	if node.Left == nil && node.Right == nil {
		fmt.Println("ender_retree=>", node.Value)
		return
	}
}

func prelist(node *Node, root bool, level int) {
	if node == nil {
		return
	}

	fmt.Println("中", level, node.Value)

	if node.Left != nil {
		fmt.Println("left=>", level+1, node.Left.Value)
		prelist(node.Left, false, level+1)
	}
	if node.Right != nil {
		fmt.Println("right=>", level+1, node.Right.Value)
		prelist(node.Right, false, level+1)
	}
}

func main() {
	value := []int{1, 8, 5, 7, 6, 9}
	rootNode := Node{Value: value[0]}
	for i := 1; i < len(value); i++ {
		tree(&rootNode, value[i])
	}
	fmt.Println("####################")
	prelist(&rootNode, true, 0)
	fmt.Println("####################")

	rethree(&rootNode)
	fmt.Println("///////////////////")
	prelist(&rootNode, true, 0)
	fmt.Println("//////////////////")
}
