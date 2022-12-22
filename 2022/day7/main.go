package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type node struct {
	parent   *node
	children []*node
	size     int
	name     string
}

func (n *node) GetChildWithName(name string) *node {
	for _, childNode := range n.children {
		if childNode.name == name {
			return childNode
		}
	}

	return nil
}

func (n *node) CreateChildWithName(name string) *node {

	child := &node{
		name:   name,
		parent: n,
	}

	n.children = append(n.children, child)
	return child
}

func main() {

	// Build file tree structure from input
	fileTree := readFS("input")

	// Find all of the directories with a total
	// size of at most 100000. What is the sum of the total sizes of those directories?

	printNodeTree(fileTree)
}

func printNodeTree(root *node) {
	q := make([]*node, 0, 100)
	q = append(q, root)

	for len(q) > 0 {
		n := q[0]
		fmt.Println(n)
		q = q[1:]
		q = append(q, n.children...)
	}
}

func readFS(p string) *node {

	s := input.OpenFileBuffered(p)

	root := &node{
		name: "root",
	}

	n := root

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
		} else {
			// Ls output
			switch l[0] {
			case "dir":
				// Dont need to handle dirs
				continue
			default:
				// File type node , add it to current node n
				fileName := l[1]
				fileSize, err := strconv.Atoi(l[0])
				if err != nil {
					log.Fatal(err)
				}
				c := n.GetChildWithName(fileName)
				if c == nil {
					c = n.CreateChildWithName(fileName)
					c.size = fileSize
				}
			}
		}

	}

	return root
}

func changeDir(n *node, dir string) *node {

	switch dir {
	case "..":
		if n.parent == nil {
			log.Fatal("No parent to set node to", n)
		}
		return n.parent
	default:
		// Search if dir exist in children , otherwise create it
		c := n.GetChildWithName(dir)
		if c != nil {
			return c
		}

		return n.CreateChildWithName(dir)
	}
}
