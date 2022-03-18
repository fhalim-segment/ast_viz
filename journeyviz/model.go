package journeyviz

import (
	"gitlab.com/fawad/ast-viz/astviz"
)

type Definition struct {
	Initial      string           `json:"initial"`
	States       map[string]State `json:"states"`
	ExitSettings *ExitSettings    `json: "exitSettings"`
}

type ExitSettings struct {
	ExitConditions *[]ExitCondition `json:"exitConditions"`
}

type ExitCondition = map[string]interface{}

type State struct {
	Meta StateMeta
	On   map[string]State `json:"on"`
}

type StateMeta struct {
	StateType    string         `json:"type"`
	Name         string         `json:"name"`
	Ast          *astviz.Node   `json:"ast"`
	Duration     *int64         `json:"duration"`
	Destinations *[]Destination `json:"destinations"`
}

type Destination struct {
	Id              string `json:"id"`
	MetadataId      string `json:"metadataId"`
	Name            string `json:"name"`
	DestinationType string `json:"type"`
}
