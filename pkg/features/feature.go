package features

type Feature interface {
	Execute()
}

type BaseFeature struct {
	Context map[string]interface{} //TODO I need to change this for something more significant.
}

func (f *BaseFeature) GetContext() {
	//TODO See what should I do here.
}
