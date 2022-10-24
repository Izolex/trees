package src

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"strings"
)

type AvlNode[K constraints.Ordered] struct {
	left, right *AvlNode[K]
	height      int
	key         K
}

type AvlTree[K constraints.Ordered] struct {
	root *AvlNode[K]
}

func NewAvlTree[K constraints.Ordered]() *AvlTree[K] {
	return &AvlTree[K]{}
}

func (tree *AvlTree[K]) Insert(key K) *AvlNode[K] {
	tree.root = tree.insertRecursive(tree.root, key)
	return nil
}

func (tree *AvlTree[K]) insertRecursive(node *AvlNode[K], key K) *AvlNode[K] {
	if node == nil {
		return &AvlNode[K]{
			height: 1,
			key:    key,
		}
	} else if key < node.key {
		node.left = tree.insertRecursive(node.left, key)
	} else {
		node.right = tree.insertRecursive(node.right, key)
	}

	tree.updateHeight(node)

	return tree.balance(node)
}

func (tree *AvlTree[K]) findParent(key K) *AvlNode[K] {
	r := tree.root
	var parent *AvlNode[K]

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

func (tree *AvlTree[K]) insertNode(parent, node *AvlNode[K]) {
	if node.key < parent.key {
		parent.left = node
	} else {
		parent.right = node
	}
}

func (tree *AvlTree[K]) calcHeight(node *AvlNode[K]) int {
	left := tree.getNodeHeight(node.left)
	right := tree.getNodeHeight(node.right)
	return int(math.Max(float64(left), float64(right))) + 1
}

func (tree *AvlTree[K]) updateHeight(node *AvlNode[K]) {
	node.height = tree.calcHeight(node)
}

func (tree *AvlTree[K]) getNodeHeight(node *AvlNode[K]) int {
	if node != nil {
		return node.height
	} else {
		return 0
	}
}

func (tree *AvlTree[K]) getBalanceFactor(node *AvlNode[K]) int {
	return tree.getNodeHeight(node.right) - tree.getNodeHeight(node.left)
}

func (tree *AvlTree[K]) balance(node *AvlNode[K]) *AvlNode[K] {
	balanceFactor := tree.getBalanceFactor(node)

	if balanceFactor < -1 {
		if tree.getBalanceFactor(node.left) <= 0 {
			node = tree.rotateRight(node)
		} else {
			node.left = tree.rotateLeft(node.left)
			node = tree.rotateRight(node)
		}
	}

	if balanceFactor > 1 {
		if tree.getBalanceFactor(node.right) >= 0 {
			node = tree.rotateLeft(node)
		} else {
			node.right = tree.rotateRight(node.right)
			node = tree.rotateLeft(node)
		}
	}

	return node
}

func (tree *AvlTree[K]) rotateLeft(node *AvlNode[K]) *AvlNode[K] {
	right := node.right

	node.right = right.left
	right.left = node

	tree.updateHeight(node)
	tree.updateHeight(right)
	tree.updateRoot(node, right)

	return right
}

func (tree *AvlTree[K]) rotateRight(node *AvlNode[K]) *AvlNode[K] {
	left := node.left

	node.left = left.right
	left.right = node

	tree.updateHeight(node)
	tree.updateHeight(left)
	tree.updateRoot(node, left)

	return left
}

func (tree *AvlTree[K]) updateRoot(node, newRoot *AvlNode[K]) {
	if tree.root == node {
		tree.root = newRoot
	}
}

func (tree *AvlTree[K]) Print() {
	if tree.root == nil {
		fmt.Println("empty")
	} else {
		tree.printNode(tree.root, 1, "root")
	}
	fmt.Println()
}

func (tree *AvlTree[K]) printNode(node *AvlNode[K], level int, mark string) {
	fmt.Printf(
		"%s %s %d %d\n",
		strings.Repeat("-", level),
		mark,
		node.key,
		node.height,
	)
	if node.left != nil {
		tree.printNode(node.left, level+1, "L")
	}
	if node.right != nil {
		tree.printNode(node.right, level+1, "R")
	}
}
