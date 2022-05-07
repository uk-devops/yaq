package input

import (
	"errors"
	"strings"

	"github.com/saliceti/yaq/internal/pipeline"
)

type stringOutputFunctionType func(string) (string, error)
type mapOutputFunctionType func(string) (pipeline.StructuredData, error)

type inputFunc struct {
	stringOutputFunction stringOutputFunctionType
	mapOutputFunction    mapOutputFunctionType
}

type inputFunctionRegister map[string]inputFunc

var register = inputFunctionRegister{}

func RegisterStringFunction(name string, inputFunction stringOutputFunctionType) {
	register[name] = inputFunc{stringOutputFunction: inputFunction}
}
func RegisterMapFunction(name string, inputFunction mapOutputFunctionType) {
	register[name] = inputFunc{mapOutputFunction: inputFunction}
}

func ReadString(inputArg string) (string, error) {
	inputArgArray := strings.Split(inputArg, ":")
	if f, ok := register[inputArgArray[0]]; ok {
		var parameter string
		if len(inputArgArray) == 1 {
			parameter = ""
		} else {
			parameter = inputArgArray[1]
		}
		inputString, err := f.stringOutputFunction(parameter)
		if err != nil {
			return inputString, err
		}
		return inputString, nil
	}

	return "", errors.New("Unknown input: " + inputArg)
}

func ReadMap(inputArg string) (pipeline.StructuredData, error) {
	inputArgArray := strings.Split(inputArg, ":")
	if f, ok := register[inputArgArray[0]]; ok {
		var parameter string
		if len(inputArgArray) == 1 {
			parameter = ""
		} else {
			parameter = inputArgArray[1]
		}
		inputMap, err := f.mapOutputFunction(parameter)
		if err != nil {
			return inputMap, err
		}
		return inputMap, nil
	}

	return nil, errors.New("Unknown input: " + inputArg)
}
