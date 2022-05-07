package transform_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/saliceti/yaq/internal/pipeline"
	"github.com/saliceti/yaq/internal/transform"
)

var _ = Describe("jq", func() {
	var output pipeline.StructuredData
	var err error
	mapInput := pipeline.GenericMap{"a": 1, "b": "c"}

	Context("map input", func() {

		Context("single output", func() {
			Context("map output", func() {
				BeforeEach(func() {
					jqExpression := ". | with_entries(.key |= \"TF_VAR_\" + .)"
					output, err = transform.ProcessWithJQ(mapInput, jqExpression)
					Expect(err).NotTo(HaveOccurred())
				})
				It("produces a map", func() {
					Expect(output).To(Equal(pipeline.GenericMap{"TF_VAR_a": 1, "TF_VAR_b": "c"}))
				})
			})

			Context("array output", func() {
				BeforeEach(func() {
					jqExpression := "keys"
					output, err = transform.ProcessWithJQ(mapInput, jqExpression)
					Expect(err).NotTo(HaveOccurred())
				})
				It("produces an array", func() {
					Expect(output).To(Equal(pipeline.GenericArray{"a", "b"}))
				})
			})

			Context("string output", func() {
				BeforeEach(func() {
					jqExpression := ".b"
					output, err = transform.ProcessWithJQ(mapInput, jqExpression)
					Expect(err).NotTo(HaveOccurred())
				})
				It("produces a single key map", func() {
					Expect(output).To(Equal(pipeline.GenericMap{"result": "c"}))
				})
			})

			Context("int output", func() {
				BeforeEach(func() {
					jqExpression := ".a"
					output, err = transform.ProcessWithJQ(mapInput, jqExpression)
					Expect(err).NotTo(HaveOccurred())
				})
				It("produces a single key map", func() {
					Expect(output).To(Equal(pipeline.GenericMap{"result": 1}))
				})
			})

			Context("bool output", func() {
				BeforeEach(func() {
					jqExpression := "has(\"a\")"
					output, err = transform.ProcessWithJQ(mapInput, jqExpression)
					Expect(err).NotTo(HaveOccurred())
				})
				It("produces a single key map", func() {
					Expect(output).To(Equal(pipeline.GenericMap{"result": true}))
				})
			})
		})

		Context("multiple outputs", func() {
			BeforeEach(func() {
				jqExpression := "keys | .[]"
				output, err = transform.ProcessWithJQ(mapInput, jqExpression)
				Expect(err).NotTo(HaveOccurred())
			})
			It("produces an array", func() {
				Expect(output).To(Equal(pipeline.GenericArray{"a", "b"}))
			})

		})
	})

})
