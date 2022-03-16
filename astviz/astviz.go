package astviz

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func Load(data []byte) (*Node, error) {
	var node Node
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (n *Node) Dump(o io.StringWriter, indent int) {
	if n == nil {
		return
	}
	write(o, "(", indent)
	write(o, fmt.Sprintf("%s \"%s\"", *n.NodeType, *n.Value), 0)
	for i, c := range n.Children {
		if i > 0 {
			write(o, ", ", 0)
		}
		write(o, "\n", 0)
		c.Dump(o, indent+1)
	}
	if n.Options != nil {
		write(o, "\n", 0)
		write(o, "('options {", indent+1)
		for k, v := range n.Options {
			write(o, fmt.Sprintf("%v: %v, ", k, v), 0)
		}
		write(o, ")", 0)
	}
	write(o, ")", 0)
}
func write(o io.StringWriter, s string, indent int) {
	indentChar := "\t"
	o.WriteString(strings.Repeat(indentChar, indent) + s)
}
