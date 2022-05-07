package output

import (
	"errors"

	"github.com/saliceti/yaq/internal/pipeline"
)

type stringFunc func(string, string) error
type mapFunc func(pipeline.StructuredData, string, []string) error

type outputFunc struct {
	stringFunc stringFunc
	mapFunc    mapFunc
}
type outputFuncRegister map[string]outputFunc

var register = outputFuncRegister{}

func (f outputFunc) isMapFunc() bool {
	return f.mapFunc != nil
}

func RegisterStringFunction(name string, outputFunction stringFunc) {
	register[name] = outputFunc{stringFunc: outputFunction}
}
func RegisterMapFunction(name string, outputFunction mapFunc) {
	register[name] = outputFunc{mapFunc: outputFunction}
}

func PushString(outputArg string, outputString string) error {
	outputName, parameter := pipeline.SplitArg(outputArg)

	f, err := lookupOutputFunc(outputName)
	if err != nil {
		return err
	}

	return f.stringFunc(outputString, parameter)
}

func PushMap(outputArg string, outputData pipeline.StructuredData, extra []string) error {
	outputName, parameter := pipeline.SplitArg(outputArg)

	f, err := lookupOutputFunc(outputName)
	if err != nil {
		return err
	}

	return f.mapFunc(outputData, parameter, extra)
}

func RequiresMap(outputArg string) (bool, error) {
	outputName, _ := pipeline.SplitArg(outputArg)

	f, err := lookupOutputFunc(outputName)
	if err != nil {
		return false, err
	}

	return f.isMapFunc(), nil
}

func lookupOutputFunc(outputName string) (*outputFunc, error) {
	if f, ok := register[outputName]; ok {
		return &f, nil
	}
	return nil, errors.New("Unknown output: " + outputName)
}
