package load_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"strconv"

	. "github.com/saliceti/yaq/internal/load"
	"github.com/saliceti/yaq/internal/pipeline"
)

var _ = Describe("json", func() {
	Context("json is valid", func() {
		It("generates a map", func() {
			output, err := Unmarshal(`{
"a": "1",
"b": {
    "c": 2
  }
}`)
			// Retrieve Map
			m1, ok := output.(pipeline.GenericMap)
			Expect(ok).To(BeTrue())

			// Unmarshalled as int
			Expect(m1["a"]).To(Equal(strconv.Itoa(1)))
			// Converted to map[string]interface{}
			value := m1["b"].(map[string]interface{})["c"]
			// Typed as float
			Expect(value).To(Equal(2.0))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("jaon array", func() {
		It("generates an array", func() {
			output, err := Unmarshal(`[
  "a",
  {"b": 1}
]`)

			Expect(err).NotTo(HaveOccurred())
			// Retrieve Map
			a, ok := output.(pipeline.GenericArray)
			Expect(ok).To(BeTrue())

			Expect(a[0]).To(Equal("a"))

			// Converted to map[string]interface{} instead of pipeline.GenericMap
			m, ok := a[1].(map[string]interface{})
			Expect(ok).To(BeTrue())
			// Typed as int
			i, ok := m["b"].(float64)
			Expect(ok).To(BeTrue())
			Expect(i).To(Equal(1.0))
		})
	})

})
