package dump_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/saliceti/yaq/dump"

	"github.com/saliceti/yaq/pipeline"
)

func testDump(m interface{}) (string, error) {
	return "{structured output}", nil
}

var _ = Describe("Register", func() {
	Context("real function: json", func() {
		It("registers json", func() {
			Expect(reflect.ValueOf(FunctionRegister["json"]).Pointer()).To(
				Equal(reflect.ValueOf(DumpToJSON).Pointer()))
		})
	})
	Context("the function exists", func() {
		It("registers the function", func() {
			Register("t1", testDump)
			Expect(reflect.ValueOf(FunctionRegister["t1"]).Pointer()).To(
				Equal(reflect.ValueOf(testDump).Pointer()))
		})
		It("the function is called successfully", func() {
			Register("t2", testDump)
			Expect(FunctionRegister["t2"](pipeline.GenericMap{})).To(Equal("{structured output}"))
		})
	})
	Context("the function does not exist", func() {
		It("throws an error", func() {
			output, err := MapToString("t3", pipeline.GenericMap{})
			Expect(err).To(MatchError("Unknown dump format: t3"))
			Expect(output).To(Equal(""))
		})
	})
})
