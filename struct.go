package xml2json

import (
	"strings"
)

// Node is a data element on a tree
type Node struct {
	Children              map[string]Nodes
	Data                  string
	ChildrenAlwaysAsArray bool
	Prefix                string
	ChildrenKeys          []string
}

type Nodes []*Node

func (n *Node) AddChild(s string, c *Node) {
	if n.Children == nil {
		n.Children = map[string]Nodes{}
	}

	if _, exists := n.Children[s]; !exists {
		n.ChildrenKeys = append(n.ChildrenKeys, s)
	}

	n.Children[s] = append(n.Children[s], c)
}

func (n *Node) IsComplex() bool {
	return len(n.Children) > 0
}

func (n *Node) AddNamespacePrefix(prefix string) {
	n.Prefix = prefix
}

func (n *Node) GetChild(path string) *Node {
	result := n
	names := strings.Split(path, ".")
	for _, name := range names {
		children, exists := result.Children[name]
		if !exists {
			return nil
		}
		if len(children) == 0 {
			return nil
		}
		result = children[0]
	}
	return result
}
