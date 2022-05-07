package dump_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/saliceti/yaq/internal/dump"

	"github.com/saliceti/yaq/internal/pipeline"
)

var _ = Describe("json", func() {
	It("generates a map", func() {
		subMap := map[string]int{"c": 2}
		inputMap := pipeline.GenericMap{
			"a": "1",
			"b": subMap,
		}
		outputString, err := MapToString("json", inputMap)
		Expect(err).NotTo(HaveOccurred())
		Expect(outputString).To(Equal(`{
  "a": "1",
  "b": {
    "c": 2
  }
}`))
	})
})
