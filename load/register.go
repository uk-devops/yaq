package load

import (
	"errors"
)

type loadFunctionType func(string) (interface{}, error)
type loadFunctionRegister map[string]loadFunctionType

var FunctionRegister = loadFunctionRegister{}

func Register(name string, loadFunction loadFunctionType) {
	FunctionRegister[name] = loadFunction
}

func Unmarshal(inputString string) (interface{}, error) {
	var d interface{}
	var err error

	d, err = FunctionRegister["json"](inputString)
	if err != nil {
		d, err = FunctionRegister["yaml"](inputString)
	}

	if err != nil {
		return d, errors.New("Invalid json or yaml:\n" + err.Error())
	}

	return d, err
}
