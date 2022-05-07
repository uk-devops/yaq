package input

import (
	"errors"

	"github.com/saliceti/yaq/internal/pipeline"
)

type stringFunc func(string) (string, error)
type mapFunc func(string) (pipeline.StructuredData, error)

type inputFunc struct {
	stringFunc stringFunc
	mapFunc    mapFunc
}

type inputFuncRegister map[string]inputFunc

var register = inputFuncRegister{}

func RegisterStringFunction(name string, inputFunction stringFunc) {
	register[name] = inputFunc{stringFunc: inputFunction}
}
func RegisterMapFunction(name string, inputFunction mapFunc) {
	register[name] = inputFunc{mapFunc: inputFunction}
}

func (f inputFunc) isMapFunc() bool {
	return f.mapFunc != nil
}

func ReadString(inputArg string) (string, error) {
	inputName, parameter := pipeline.SplitArg(inputArg)

	f, err := lookupInputFunc(inputName)
	if err != nil {
		return "", err
	}

	return f.stringFunc(parameter)
}

func ReadMap(inputArg string) (pipeline.StructuredData, error) {
	inputName, parameter := pipeline.SplitArg(inputArg)

	f, err := lookupInputFunc(inputName)
	if err != nil {
		return nil, err
	}

	return f.mapFunc(parameter)
}

func CreatesMap(inputArg string) (bool, error) {
	inputName, _ := pipeline.SplitArg(inputArg)

	f, err := lookupInputFunc(inputName)
	if err != nil {
		return false, err
	}

	return f.isMapFunc(), nil
}

func lookupInputFunc(inputName string) (*inputFunc, error) {
	if f, ok := register[inputName]; ok {
		return &f, nil
	}
	return nil, errors.New("Unknown input: " + inputName)
}
