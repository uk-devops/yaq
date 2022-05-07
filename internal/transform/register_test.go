package transform_test

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/saliceti/yaq/internal/pipeline"
	. "github.com/saliceti/yaq/internal/transform"
)

func testTransform(sd pipeline.StructuredData, arg string) (pipeline.StructuredData, error) {
	return sd, nil
}
func testTransformWithArg(sd pipeline.StructuredData, arg string) (pipeline.StructuredData, error) {
	mapData := sd.(pipeline.GenericMap)
	mapData["a"] = arg
	return mapData, nil
}

var _ = Describe("Register", func() {
	Context("the function exists", func() {
		It("registers the function", func() {
			RegisterTransformFunction("t1", testTransform)
			Expect(reflect.ValueOf(FunctionRegister["t1"]).Pointer()).To(
				Equal(reflect.ValueOf(testTransform).Pointer()))
		})
		It("the function is called successfully", func() {
			RegisterTransformFunction("t2", testTransform)
			Expect(TransformWith("t2", pipeline.GenericMap{})).To(Equal(pipeline.GenericMap{}))
		})
	})
	Context("function with argument", func() {
		It("the function is called successfully", func() {
			RegisterTransformFunction("t3", testTransformWithArg)
			Expect(TransformWith("t3:something", pipeline.GenericMap{})).To(Equal(pipeline.GenericMap{"a": "something"}))
		})
	})
})
