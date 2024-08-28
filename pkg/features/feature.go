package features

import "github.com/nanicienta/nani-commons/pkg/model"

type Feature interface {
	Execute(model.Node, model.Workflow, map[string]interface{}) (
		map[string]interface{}, string,
		error,
	)
	GetInternalName() string
}

type BaseFeature struct {
	Context map[string]interface{} //TODO I need to change this for something more significant.
}

func (f *BaseFeature) GetContext() {
	//TODO See what should I do here.
}

func (f *BaseFeature) GetInternalName() string {
	return "BaseFeature"
}
