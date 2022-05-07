package input_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"reflect"

	. "github.com/saliceti/yaq/internal/input"
)

func testInputNoParam(inputParameter string) (string, error) {
	return "the output", nil
}

func testInputWithParam(inputParameter string) (string, error) {
	return inputParameter, nil
}

var _ = Describe("Register", func() {
	Context("real function: stdin", func() {
		It("registers stdin", func() {
			Expect(reflect.ValueOf(FunctionRegister["stdin"]).Pointer()).To(
				Equal(reflect.ValueOf(ReadFromStdin).Pointer()))
		})
	})
	Context("the function exists", func() {
		It("registers the function", func() {
			Register("t1", testInputNoParam)
			Expect(reflect.ValueOf(FunctionRegister["t1"]).Pointer()).To(
				Equal(reflect.ValueOf(testInputNoParam).Pointer()))
		})
		Context("no parameter", func() {
			It("the function is called successfully", func() {
				Register("t2", testInputNoParam)
				output, err := ReadString("t2")
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(Equal("the output"))
			})
		})
		Context("a parameter is passed", func() {
			It("calls the function with the parameter", func() {
				Register("t4", testInputWithParam)
				output, err := ReadString("t4:the-parameter")
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(Equal("the-parameter"))
			})
		})
		Context("2 parameters are passed", func() {
			It("ignores the second parameter", func() {
				Register("t5", testInputWithParam)
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

})
