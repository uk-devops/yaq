package input

import (
	"errors"
	"strings"
)

type inputFunctionType func(string) (string, error)
type inputFunctionRegister map[string]inputFunctionType

var FunctionRegister = inputFunctionRegister{}

func Register(name string, inputFunction inputFunctionType) {
	FunctionRegister[name] = inputFunction
}

func ReadString(inputArg string) (string, error) {
	inputArgArray := strings.Split(inputArg, ":")
	if f, ok := FunctionRegister[inputArgArray[0]]; ok {
		var parameter string
		if len(inputArgArray) == 1 {
			parameter = ""
		} else {
			parameter = inputArgArray[1]
		}
		inputString, err := f(parameter)
		if err != nil {
			return inputString, err
		}
		return inputString, nil
	}

	return "", errors.New("Unknown input: " + inputArg)
}
