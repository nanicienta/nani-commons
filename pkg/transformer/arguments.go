package transformer

import (
	argumentspkg "github.com/nanicienta/nani-commons/pkg/constants/arguments"
	"github.com/nanicienta/nani-commons/pkg/errors"
	"github.com/nanicienta/nani-commons/pkg/model"
)

func TransformArgumentsToQueryArguments(arguments []model.Argument, wf model.Workflow) (
	map[string]interface{},
	error,
) {
	queryArguments := make(map[string]interface{})
	for _, argument := range arguments {
		if argument.Type == argumentspkg.ArgumentTypeVariable {
			variable, found := wf.GetVariable(argument)
			if !found {
				return nil, errors.NewErrorVariableNotFound(argument.Value)
			}
			queryArguments[argument.Name] = variable
		} else if argument.Type == argumentspkg.ArgumentTypeValue {
			queryArguments[argument.Name] = argument.Value
		} else if argument.Type == argumentspkg.ArgumentTypePreviousResult {
			//TODO implement this shit
		}
	}
	return queryArguments, nil
}
