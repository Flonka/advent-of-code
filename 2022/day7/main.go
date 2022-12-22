package main

import (
	"fmt"
	"log"
	"sort"
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

	updateSizesRecursion(fileTree)

	printTree(fileTree)
}

func updateSizesRecursion(n *node) int {
	for _, child := range n.children {
		// Isdir (child slice is nil), should be a flag..
		if child.children != nil {
			n.size += updateSizesRecursion(child)
		} else {
			n.size += child.size
		}
	}
	return n.size
}

func printTree(root *node) {
	q := make([]*node, 0, 100)
	q = append(q, root)

	// Find all of the directories with a total
	// size of at most 100000. What is the sum of the total sizes of those directories?
	var part1 int

	// Part2
	// The total disk space available to the filesystem is 70000000.
	// To run the update, you need unused space of at least 30000000.
	unused := 70000000 - root.size
	needed := 30000000 - unused

	var part2 []int

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		for _, child := range n.children {
			if child.children != nil {
				q = append(q, child)
			}
		}

		// If it is a dir and size limit
		if n.size <= 100000 {
			part1 += n.size
		}

		if n.size >= needed {
			part2 = append(part2, n.size)
		}
	}
	sort.Ints(part2)
	fmt.Println("part1", part1)
	fmt.Println("part2", part2[0])

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
