package output_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/output"
	"github.com/saliceti/yaq/pipeline"
)

var iHaveBeenCalled = false

func testOutputFromString(s, param string) error {
	iHaveBeenCalled = true
	return nil
}

func testOutputFromMap(s pipeline.StructuredData, extra []string) error {
	iHaveBeenCalled = true
	return nil
}

var _ = Describe("Register", func() {
	Context("real function: stdout", func() {
		It("registers stdout", func() {
			Expect(reflect.ValueOf(StringFunctionRegister["stdout"]).Pointer()).To(
				Equal(reflect.ValueOf(PushToStdout).Pointer()))
		})
	})
	Context("the string function exists", func() {
		It("registers the string function", func() {
			RegisterStringFunction("t1", testOutputFromString)
			Expect(reflect.ValueOf(StringFunctionRegister["t1"]).Pointer()).To(
				Equal(reflect.ValueOf(testOutputFromString).Pointer()))
		})
		It("the string function is called successfully", func() {
			RegisterStringFunction("t2", testOutputFromString)
			StringFunctionRegister["t2"]("dummy", "parameter")
			Expect(iHaveBeenCalled).To(BeTrue())
		})
	})
	Context("the string function does not exist", func() {
		It("throws an error", func() {
			err := PushString("t3", "dummy")
			Expect(err).To(MatchError("Unknown output: t3"))
		})
	})

	Context("the map function exists", func() {
		It("registers the map function", func() {
			RegisterMapFunction("t4", testOutputFromMap)
			Expect(reflect.ValueOf(MapFunctionRegister["t4"]).Pointer()).To(
				Equal(reflect.ValueOf(testOutputFromMap).Pointer()))
		})
		It("the map function is called successfully", func() {
			RegisterMapFunction("t5", testOutputFromMap)
			MapFunctionRegister["t5"](pipeline.GenericMap{"dummy": "value"}, nil)
			Expect(iHaveBeenCalled).To(BeTrue())
		})
	})
	Context("the map function does not exist", func() {
		It("throws an error", func() {
			err := PushMap("t6", pipeline.GenericMap{"dummy": "value"}, nil)
			Expect(err).To(MatchError("Unknown output: t6"))
		})
	})
})
