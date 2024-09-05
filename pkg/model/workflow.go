package model

import "github.com/nanicienta/nani-commons/pkg/model/connections"

type Workflow struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Init        string `json:"init"`
	//TODO como mierda hago esto, no es lo mismo esta connection que hacer algo como un cliente rest
	Connections        []connections.BaseConnection `json:"connections"`
	Nodes              []Node                       `json:"nodes"`
	indexedNodes       map[string]Node
	indexedConnections map[string]connections.BaseConnection
	Context            map[string]interface{}
}

func (w *Workflow) GetNode(id string) (bool, Node) {
	node, exists := w.indexedNodes[id]
	return exists, node
}

func (w *Workflow) InitWorkflow() {
	w.indexedNodes = make(map[string]Node)
	w.Context = make(map[string]interface{})
	w.indexedConnections = make(map[string]connections.BaseConnection)
	for _, node := range w.Nodes {
		node.InitNode()
		w.indexedNodes[node.Id] = node
	}
	for _, conn := range w.Connections {
		w.indexedConnections[conn.Id] = conn
	}
}

func (w *Workflow) GetConnection(connectionId string) (bool, connections.BaseConnection) {
	connection, exists := w.indexedConnections[connectionId]
	return exists, connection
}

func (w *Workflow) GetVariable(argument Argument) (interface{}, bool) {
	//TODO implementar
	panic("implement me")
	return nil, false
}
