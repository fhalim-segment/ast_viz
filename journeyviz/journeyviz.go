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
	destsToSources := make(map[Destination][]string, 0)
	fmt.Fprintln(o, fmt.Sprintf("@startuml\nhide empty description\n[*] --> %s", n.Initial))
	statesToWalk := []string{n.Initial}
	for len(statesToWalk) > 0 {
		stateName := statesToWalk[0]
		statesToWalk = statesToWalk[1:]
		state := n.States[stateName]
		stateTitle := stateName
		fmt.Fprintln(o, fmt.Sprintf("state \"%s\" as %s", stateTitle, stateName))
		fmt.Fprintln(o, fmt.Sprintf("%s: type:  %s", stateName, state.Meta.StateType))
		if state.Meta.Name != "" {
			fmt.Fprintln(o, fmt.Sprintf("%s: name:  %s", stateName, state.Meta.Name))
		}
		if state.Meta.StateType == "delay" && state.Meta.Duration != nil {
			fmt.Fprintln(o, fmt.Sprintf("%s: duration:  %d", stateName, *state.Meta.Duration))
		}

		if state.Meta.Ast != nil {
			s := ""
			buf := bytes.NewBufferString(s)
			state.Meta.Ast.Dump(buf, 0)
			astMessage := buf.String()

			if len(strings.Trim(state.Meta.Name, " ")) > 0 {
				stateTitle = state.Meta.Name
			}

			lines := strings.Split(astMessage, "\n")
			for _, line := range lines {
				fmt.Fprintln(o, fmt.Sprintf("%s: |  %s", stateName, line))
			}
		}
		if state.On != nil {
			for k, _ := range state.On {
				if n.States[k].Meta.StateType == "placeholder" {
					continue
				}
				fmt.Fprintln(o, fmt.Sprintf("%s --> %s", stateName, k))
				statesToWalk = append(statesToWalk, k)
			}
		}

		if state.Meta.Destinations != nil {
			for _, dest := range *state.Meta.Destinations {
				sources := destsToSources[dest]
				destsToSources[dest] = append(sources, stateName)
			}
		}
	}
	if len(destsToSources) > 0 {
		fmt.Fprintln(o, "state Destinations {")
		for dest, _ := range destsToSources {
			fmt.Fprintln(o, fmt.Sprintf("  state \"%s\" as %s", dest.Name, dest.Id))
			fmt.Fprintln(o, fmt.Sprintf("  %s: Type: %s", dest.Id, dest.DestinationType))
			fmt.Fprintln(o, fmt.Sprintf("  %s: MetadataId: %s", dest.Id, dest.MetadataId))
		}
		fmt.Fprintln(o, "}")
		for dest, sources := range destsToSources {
			for _, source := range sources {
				fmt.Fprintln(o, fmt.Sprintf("%s --> %s", source, dest.Id))
			}
		}
	}
	fmt.Fprintln(o, "@enduml")
}
func escapeForState(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\"", "\\\"\""), "\n", "\\n")
}
