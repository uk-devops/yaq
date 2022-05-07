package load_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/load"
	"github.com/saliceti/yaq/pipeline"
)

func testLoad(dummy string) (pipeline.GenericMap, error) {
	return pipeline.GenericMap{}, nil
}

var _ = Describe("Register", func() {
	Context("real function: json", func() {
		It("registers json", func() {
			Expect(reflect.ValueOf(FunctionRegister["json"]).Pointer()).To(
				Equal(reflect.ValueOf(LoadFromJSON).Pointer()))
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
			outputMap, err := MapFromString(`not structured`)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid json or yaml")))
			Expect(outputMap).To(Equal(pipeline.GenericMap{}))
		})
	})
})
