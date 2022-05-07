package load_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/load"
	"github.com/saliceti/yaq/pipeline"
)

var _ = Describe("yaml", func() {
	Context("yaml map", func() {
		It("generates a map", func() {
			output, err := Unmarshal(`---
a: 1
b:
  c: 2
`)
			Expect(err).NotTo(HaveOccurred())
			// Retrieve Map
			m1, ok := output.(pipeline.GenericMap)
			Expect(ok).To(BeTrue())
			// Unmarshalled as int
			Expect(m1["a"]).To(Equal(1))
			// Converted to pipeline.GenericMap
			m2, ok := m1["b"].(pipeline.GenericMap)
			Expect(ok).To(BeTrue())
			// Typed as int
			i, ok := m2["c"].(int)
			Expect(ok).To(BeTrue())
			Expect(i).To(Equal(2))
		})
	})

	Context("yaml list", func() {
		It("generates an array", func() {
			output, err := Unmarshal(`---
- a
- b: 1
`)

			Expect(err).NotTo(HaveOccurred())
			// Retrieve Array
			a, ok := output.(pipeline.GenericArray)
			Expect(ok).To(BeTrue())

			Expect(a[0]).To(Equal("a"))

			// Converted to map[string]interface{} instead of pipeline.GenericMap
			m, ok := a[1].(map[string]interface{})
			Expect(ok).To(BeTrue())
			// Typed as int
			i, ok := m["b"].(int)
			Expect(ok).To(BeTrue())
			Expect(i).To(Equal(1))
		})
	})
})
