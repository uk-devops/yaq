package output

import (
	"errors"
)

type stringOutputFunctionType func(string)
type mapOutputFunctionType func(interface{}, []string) error

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
	if f, ok := StringFunctionRegister[outputArg]; ok {
		f(outputString)
		return nil
	}
	return errors.New("Unknown output: " + outputArg)
}

func PushMap(outputArg string, outputData interface{}, extra []string) error {
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
