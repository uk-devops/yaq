package input

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/saliceti/yaq/internal/pipeline"

	"reflect"
)

func testInputNoParam(inputParameter string) (string, error) {
	return "the output", nil
}

func testInputWithParam(inputParameter string) (string, error) {
	return inputParameter, nil
}

func testInputMap(inputParameter string) (pipeline.StructuredData, error) {
	return pipeline.GenericMap{}, nil
}

var _ = Describe("Register", func() {
	Context("real function: stdin", func() {
		It("registers stdin", func() {
			Expect(reflect.ValueOf(register["stdin"].stringFunc).Pointer()).To(
				Equal(reflect.ValueOf(ReadFromStdin).Pointer()))
		})
	})
	Context("the function exists", func() {
		It("registers the function", func() {
			RegisterStringFunction("t1", testInputNoParam)
			Expect(reflect.ValueOf(register["t1"].stringFunc).Pointer()).To(
				Equal(reflect.ValueOf(testInputNoParam).Pointer()))
		})
		Context("no parameter", func() {
			It("the function is called successfully", func() {
				RegisterStringFunction("t2", testInputNoParam)
				output, err := ReadString("t2")
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(Equal("the output"))
			})
		})
		Context("a parameter is passed", func() {
			It("calls the function with the parameter", func() {
				RegisterStringFunction("t4", testInputWithParam)
				output, err := ReadString("t4:the-parameter")
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(Equal("the-parameter"))
			})
		})
		Context("2 parameters are passed", func() {
			It("ignores the second parameter", func() {
				RegisterStringFunction("t5", testInputWithParam)
				output, err := ReadString("t5:the-parameter:error-parameter")
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(Equal("the-parameter"))
			})
		})
	})
	Context("the function does not exist", func() {
		It("throws an error", func() {
			output, err := ReadString("t3")
			Expect(err).To(MatchError("Unknown input: t3"))
			Expect(output).To(Equal(""))
		})
	})
	Context("map function", func() {
		It("registers the function", func() {
			RegisterMapFunction("t6", testInputMap)
			Expect(reflect.ValueOf(register["t6"].mapFunc).Pointer()).To(
				Equal(reflect.ValueOf(testInputMap).Pointer()))
		})
	})
})
