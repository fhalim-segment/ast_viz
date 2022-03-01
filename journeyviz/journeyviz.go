package journeyviz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Load(data []byte) (*Definition, error) {
	var journey Definition
	err := json.Unmarshal(data, &journey)
	if err != nil {
		return nil, err
	}
	return &journey, nil
}

func (n *Definition) Dump(o *os.File, indent int) {
	fmt.Fprintln(o, fmt.Sprintf("@startuml\nhide empty description\n[*] -down-> %s", n.Initial))
	statesToWalk := []string{n.Initial}
	for len(statesToWalk) > 0 {
		stateName := statesToWalk[0]
		statesToWalk = statesToWalk[1:]
		state := n.States[stateName]
		if state.Meta.Ast != nil {
			s := ""
			buf := bytes.NewBufferString(s)
			state.Meta.Ast.Dump(buf, 0)
			astMessage := buf.String()
			stateTitle := stateName
			if len(strings.Trim(state.Meta.Name, " ")) > 0 {
				stateTitle = state.Meta.Name
			}
			fmt.Fprintln(o, fmt.Sprintf("state \"%s\" as %s", stateTitle, stateName))
			lines := strings.Split(astMessage, "\n")
			for _, line := range lines {
				fmt.Fprintln(o, fmt.Sprintf("%s: |  %s", stateName, line))
			}
		}
		if state.On != nil {
			for k, _ := range state.On {
				fmt.Fprintln(o, fmt.Sprintf("%s -down-> %s", stateName, k))
				statesToWalk = append(statesToWalk, k)
			}
		}
	}
	fmt.Fprintln(o, "@enduml")
}
func escapeForState(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\"", "\\\"\""), "\n", "\\n")
}
