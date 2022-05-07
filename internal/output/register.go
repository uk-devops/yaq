package output

import (
	"errors"
	"strings"

	"github.com/saliceti/yaq/internal/pipeline"
)

type stringOutputFunctionType func(string, string) error
type mapOutputFunctionType func(pipeline.StructuredData, []string) error

type outputFunctionRegister map[string]stringOutputFunctionType
type mapOutputFunctionRegister map[string]mapOutputFunctionType

var StringFunctionRegister = outputFunctionRegister{}
var MapFunctionRegister = mapOutputFunctionRegister{}

func RegisterStringFunction(name string, outputFunction stringOutputFunctionType) {
	StringFunctionRegister[name] = outputFunction
}
func RegisterMapFunction(name string, outputFunction mapOutputFunctionType) {
	MapFunctionRegister[name] = outputFunction
}

func PushString(outputArg string, outputString string) error {
	outputArgArray := strings.Split(outputArg, ":")

	if f, ok := StringFunctionRegister[outputArgArray[0]]; ok {
		var parameter string
		if len(outputArgArray) == 1 {
			parameter = ""
		} else {
			parameter = outputArgArray[1]
		}

		err := f(outputString, parameter)
		return err
	}
	return errors.New("Unknown output: " + outputArg)
}

func PushMap(outputArg string, outputData pipeline.StructuredData, extra []string) error {
	if f, ok := MapFunctionRegister[outputArg]; ok {
		err := f(outputData, extra)
		return err
	}
	return errors.New("Unknown output: " + outputArg)
}

func RequiresMap(outputArg string) bool {
	if _, ok := MapFunctionRegister[outputArg]; ok {
		return true
	}
	return false
}
