package load

import (
	"errors"

	"github.com/saliceti/yaq/pipeline"
)

type loadFunctionType func(string) (pipeline.GenericMap, error)
type loadFunctionRegister map[string]loadFunctionType

var FunctionRegister = loadFunctionRegister{}

func Register(name string, loadFunction loadFunctionType) {
	FunctionRegister[name] = loadFunction
}

func MapFromString(inputString string) (pipeline.GenericMap, error) {
	inputMap, err := FunctionRegister["json"](inputString)
	if err != nil {
		inputMap, err = FunctionRegister["yaml"](inputString)
	}
	if err != nil {
		// log.Println(err.Error())
		return pipeline.GenericMap{}, errors.New("Invalid json or yaml:\n" + err.Error())
	}
	return inputMap, err
}
