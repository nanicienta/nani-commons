package model

import "github.com/nanicienta/nani-commons/pkg/model/connections"

type Workflow struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Init        string `json:"init"`
	//TODO como mierda hago esto, no es lo mismo esta connection que hacer algo como un cliente rest
	Connections  []connections.BaseConnection `json:"connections"`
	Nodes        []Node                       `json:"nodes"`
	indexedNodes map[string]Node
	indexedLink  map[string]Link
}
