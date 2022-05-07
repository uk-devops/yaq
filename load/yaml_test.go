package load_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/load"
	"github.com/saliceti/yaq/pipeline"
)

var _ = Describe("yaml", func() {
	Context("yaml is valid", func() {
		It("generates a map", func() {
			outputMap, err := MapFromString(`---
a: 1
b:
  c: 2
`)
			// Unmarshalled as int
			Expect(outputMap["a"]).To(Equal(1))
			// Converted to pipeline.GenericMap
			value := outputMap["b"].(pipeline.GenericMap)["c"]
			// Typed as int
			Expect(value).To(Equal(2))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
