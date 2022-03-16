package astviz

type Node struct {
	Children []Node                 `json:"children"`
	NodeType *string                `json:"type"`
	Value    *string                `json:"value"`
	Options  map[string]interface{} `json:"options"`
}
