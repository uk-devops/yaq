package input

import (
	"errors"

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
	inputName, parameter := pipeline.SplitArg(inputArg)

	f, err := lookupInputFunction(inputName)
	if err != nil {
		return "", err
	}

	return f.stringOutputFunction(parameter)
}

func ReadMap(inputArg string) (pipeline.StructuredData, error) {
	inputName, parameter := pipeline.SplitArg(inputArg)

	f, err := lookupInputFunction(inputName)
	if err != nil {
		return nil, err
	}

	return f.mapOutputFunction(parameter)
}

func lookupInputFunction(inputName string) (*inputFunc, error) {
	if f, ok := register[inputName]; ok {
		return &f, nil
	}
	return nil, errors.New("Unknown input: " + inputName)
}
