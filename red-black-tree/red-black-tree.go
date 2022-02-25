package red_black_tree

import "fmt"

type RedBlackTree struct {
	Root *Node
}

type Node struct {
	p     *Node
	left  *Node
	right *Node
	value int
	color int
}

const (
	RED int = iota
	BLACK
)

func Tree_init() *RedBlackTree {
	T := new(RedBlackTree)
	return T
}

func Insert_node(T *RedBlackTree, value int) {
	// initinize for new node
	node := new(Node)
	node.value = value
	node.color = RED

	if T.Root == nil {
		T.Root = node
		return
	}
	cur := T.Root
	var parent *Node
	//find parent node of node
	for cur != nil {
		parent = cur
		if node.value < cur.value {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	//being left child
	if node.value < parent.value {
		parent.left = node
		node.p = parent
	} else { //being right child
		parent.right = node
		node.p = parent
	}
}

func Print(root *Node) {
	if root != nil {
		Print(root.left)
		fmt.Println("visited ", root.value)
		Print(root.right)
	}
}
