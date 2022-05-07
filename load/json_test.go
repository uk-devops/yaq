package load_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strconv"

	. "github.com/saliceti/yaq/load"
)

var _ = Describe("json", func() {
	Context("json is valid", func() {
		It("generates a map", func() {
			outputMap, err := MapFromString(`{
"a": "1",
"b": {
    "c": 2
  }
}`)
			// Unmarshalled as int
			Expect(outputMap["a"]).To(Equal(strconv.Itoa(1)))
			// Converted to map[string]interface{}
			value := outputMap["b"].(map[string]interface{})["c"]
			// Typed as float
			Expect(value).To(Equal(2.0))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
