package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type node struct {
	parent   *node
	children []*node
	size     int
	name     string
}

func main() {

	// Build file tree structure from input
	fileTree := readFS("input")

	fmt.Println(fileTree)
	// Find all of the directories with a total
	// size of at most 100000. What is the sum of the total sizes of those directories?
}

func readFS(p string) *node {

	s := input.OpenFileBuffered(p)

	n := &node{}
	n.children = make([]*node, 0, 5)

	var cmd string
	for s.Scan() {

		l := strings.Split(s.Text(), " ")

		if l[0] == "$" {
			// Command line
			cmd = l[1]

			switch cmd {
			case "cd":
				targetDir := l[2]
				n = changeDir(n, targetDir)
			}
		}

	}

	return n
}

func changeDir(n *node, dir string) *node {

	switch dir {
	case "..":
		if n.parent == nil {
			log.Fatal("No parent to set nodde to", n)
		}
		return n.parent
	default:
		// Search if dir exist in children , otherwise create it
		for _, childNode := range n.children {
			if childNode.name == dir {
				return childNode
			}
		}
		child := &node{
			name:     dir,
			parent:   n,
			children: make([]*node, 0, 5),
		}

		n.children = append(n.children, child)

		return child
	}
}
