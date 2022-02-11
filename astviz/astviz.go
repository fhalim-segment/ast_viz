package astviz

import (
	"encoding/json"
	"fmt"
	"os"
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

func (n *Node) Dump(o *os.File, indent int) {
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
	write(o, ")", 0)
}
func write(o *os.File, s string, indent int) {
	indentChar := "\t"
	o.WriteString(strings.Repeat(indentChar, indent) + s)
}
