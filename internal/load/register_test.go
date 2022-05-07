package load_test

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/internal/load"
	"github.com/saliceti/yaq/internal/pipeline"
)

func testLoad(dummy string) (pipeline.StructuredData, error) {
	return pipeline.GenericMap{}, nil
}

var _ = Describe("Register", func() {
	Context("real function: json", func() {
		It("registers json", func() {
			Expect(reflect.ValueOf(FunctionRegister["json"]).Pointer()).To(
				Equal(reflect.ValueOf(UnmarshalJSON).Pointer()))
		})
	})
	Context("the function exists", func() {
		It("registers the function", func() {
			Register("t1", testLoad)
			Expect(reflect.ValueOf(FunctionRegister["t1"]).Pointer()).To(
				Equal(reflect.ValueOf(testLoad).Pointer()))
		})
		It("the function is called successfully", func() {
			Register("t2", testLoad)
			Expect(FunctionRegister["t2"]("structured input")).To(Equal(pipeline.GenericMap{}))
		})
	})
})

var _ = Describe("invalid", func() {
	Context("not structured", func() {
		It("throws an error", func() {
			output, err := Unmarshal(`not structured`)
			Expect(err).To(MatchError(ContainSubstring("Invalid json or yaml")))
			Expect(output).To(BeNil())
		})
	})
})
