package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

var currentID int64 = 0

func nextID() string {
	id := strconv.FormatInt(currentID, 10)
	currentID++
	return id
}

type digrap struct {
	nodes []node
	haha  unsafe.Pointer
}

type node struct {
	id    string
	shape string
	label string
}

// Node add node to graph
func (d *digrap) Node(node node) error {
	return nil
}

// String encodes graph to graphviz
func (d *digrap) String() string {
	var sb strings.Builder
	for _, nod := range d.nodes {
		sb.WriteString(fmt.Sprintf("\"%s\" [shape=%s label=\"%s\"]\n", nod.id, nod.shape, nod.label))
	}

	return fmt.Sprintf(`
digraph {
	%s
}
`, sb.String())
}

func main() {
	var g = &digrap{nodes: []node{
		{id: nextID(), shape: "circle", label: "circle 1"},
		{id: nextID(), shape: "rectangle", label: "rectangle 1"},
	}}

	fmt.Print(g)
}
