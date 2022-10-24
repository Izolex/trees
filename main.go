package main

import (
	"trees/src"
)

func main() {
	t := src.NewAvlTree[int]()
	t.Insert(4)
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)
	t.Insert(8)

	t.Print()
}
