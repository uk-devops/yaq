package output

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/saliceti/yaq/internal/pipeline"
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
			Expect(reflect.ValueOf(register["stdout"].stringOutputFunction).Pointer()).To(
				Equal(reflect.ValueOf(PushToStdout).Pointer()))
		})
	})
	Context("the string function exists", func() {
		It("registers the string function", func() {
			RegisterStringFunction("t1", testOutputFromString)
			Expect(reflect.ValueOf(register["t1"].stringOutputFunction).Pointer()).To(
				Equal(reflect.ValueOf(testOutputFromString).Pointer()))
		})
		It("the string function is called successfully", func() {
			RegisterStringFunction("t2", testOutputFromString)
			register["t2"].stringOutputFunction("dummy", "parameter")
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
			Expect(reflect.ValueOf(register["t4"].mapOutputFunction).Pointer()).To(
				Equal(reflect.ValueOf(testOutputFromMap).Pointer()))
		})
		It("the map function is called successfully", func() {
			RegisterMapFunction("t5", testOutputFromMap)
			register["t5"].mapOutputFunction(pipeline.GenericMap{"dummy": "value"}, nil)
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
