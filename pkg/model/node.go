package model

type Node struct {
	Id                    string  `json:"id"`
	Type                  string  `json:"type"`
	ConnectionId          string  `json:"connectionId"`
	Fields                []Field `json:"fields"`
	Links                 []Link  `json:"links"`
	StoreResultInVariable bool    `json:"storeResultInVariable"`
	VariableName          string  `json:"variableName"`
	indexedLink           map[string]Link
	indexedFields         map[string]Field
}

func (n *Node) InitNode() {
	n.indexedLink = make(map[string]Link)
	for _, link := range n.Links {
		n.indexedLink[link.Id] = link
	}
	n.indexedFields = make(map[string]Field)
	for _, field := range n.Fields {
		n.indexedFields[field.Id] = field
	}
}

func (n *Node) GetField(fieldId string) (bool, Field) {
	field, exists := n.indexedFields[fieldId]
	return exists, field
}
