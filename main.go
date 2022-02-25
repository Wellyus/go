package main

import (
	rbt "github.com/Wellyus/go/red-black-tree"
)

func main() {
	T := rbt.Tree_init()
	for i, j := 4, 6; i >= 0; i, j = i-1, j+1 {
		rbt.Insert_node(T, i)
		rbt.Insert_node(T, j)
	}
	rbt.Print(T.Root)
}
