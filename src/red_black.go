package src

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

type Color int
type Direction int

const (
	Black Color = iota
	Red
)

const (
	Left Direction = iota
	Right
)

type redBlackNode[K constraints.Ordered] struct {
	parent, left, right *redBlackNode[K]
	color               Color
	key                 K
}

type RedBlackTree[K constraints.Ordered] struct {
	root *redBlackNode[K]
}

func NewRedBlackTree[K constraints.Ordered]() *RedBlackTree[K] {
	return &RedBlackTree[K]{}
}

func (tree *RedBlackTree[K]) Insert(key K) {
	parent := tree.findParent(key)
	n := &redBlackNode[K]{parent: parent, color: Red, key: key}

	if parent == nil {
		tree.setRoot(n)
		return
	}

	if n.key < parent.key {
		parent.left = n
	} else {
		parent.right = n
	}

	if parent.parent == nil {
		return
	}

	tree.fix(n)
	tree.root.color = Black
}

func (tree *RedBlackTree[K]) rotateLeft(n *redBlackNode[K]) {
	m := n.right
	n.right = m.left

	if m.left != nil {
		m.left.parent = n
	}

	m.parent = n.parent

	if n.parent == nil {
		tree.setRoot(m)
	} else if n == n.parent.left {
		n.parent.left = m
	} else {
		n.parent.right = m
	}

	m.left = n
	n.parent = m
}

func (tree *RedBlackTree[K]) rotateRight(n *redBlackNode[K]) {
	m := n.left
	n.left = m.right

	if m.right != nil {
		m.right.parent = n
	}

	m.parent = n.parent

	if n.parent == nil {
		tree.setRoot(m)
	} else if n == n.parent.right {
		n.parent.right = m
	} else {
		n.parent.left = m
	}

	m.right = n
	n.parent = m
}

func (tree *RedBlackTree[K]) findParent(key K) *redBlackNode[K] {
	r := tree.root
	var parent *redBlackNode[K]

	for r != nil {
		parent = r

		if key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}

	return parent
}

func (tree *RedBlackTree[K]) fix(n *redBlackNode[K]) {
	for n != nil && n != tree.root && n.parent.color == Red {
		if n.parent == n.parent.parent.right {
			n = tree.recolorAndRotate(n, Left)
		} else {
			n = tree.recolorAndRotate(n, Right)
		}
	}
}

func (tree *RedBlackTree[K]) recolorAndRotate(n *redBlackNode[K], direction Direction) *redBlackNode[K] {
	var m *redBlackNode[K]

	if direction == Left {
		m = n.parent.parent.left
	} else {
		m = n.parent.parent.right
	}

	if m != nil && m.color == Red {
		m.color = Black
		n.parent.color = Black
		n.parent.parent.color = Red

		return n.parent.parent
	} else {
		if direction == Left {
			if n == n.parent.left {
				n = n.parent
				tree.rotateRight(n)
			}
		} else {
			if n == n.parent.right {
				n = n.parent
				tree.rotateLeft(n)
			}
		}

		n.parent.color = Black
		n.parent.parent.color = Red
		if direction == Left {
			tree.rotateLeft(n.parent.parent)
		} else {
			tree.rotateRight(n.parent.parent)
		}

		return n
	}
}

func (tree *RedBlackTree[K]) setRoot(n *redBlackNode[K]) {
	n.color = Black
	tree.root = n
}

func (tree *RedBlackTree[K]) Print() {
	if tree.root == nil {
		fmt.Println("empty")
	} else {
		tree.printNode(tree.root, 1, "r")
	}
	fmt.Println()
}

func (tree *RedBlackTree[K]) printNode(n *redBlackNode[K], level int, mark string) {
	fmt.Printf(
		"%s %d (%s %d)\n",
		strings.Repeat("-", level),
		n.key,
		mark,
		n.color,
	)
	if n.left != nil {
		tree.printNode(n.left, level+1, "L")
	}
	if n.right != nil {
		tree.printNode(n.right, level+1, "R")
	}
}
