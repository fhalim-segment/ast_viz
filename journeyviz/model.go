package journeyviz

import (
	"gitlab.com/fawad/ast-viz/astviz"
)

type Definition struct {
	Initial string           `json:"initial"`
	States  map[string]State `json:"states"`
}

type State struct {
	Meta StateMeta
	On   map[string]State `json:"on"`
}

type StateMeta struct {
	StateType string       `json:"type"`
	Name      string       `json:"name"`
	Ast       *astviz.Node `json:"ast"`
}
