package model

type Node struct {
	Id     string  `json:"id"`
	Type   string  `json:"type"`
	Fields []Field `json:"fields"`
	Links  []Link  `json:"links"`
	IsEnd  bool    `json:"isEnd"`
}
