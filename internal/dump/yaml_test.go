package dump_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/uk-devops/yaq/internal/dump"

	"github.com/uk-devops/yaq/internal/pipeline"
)

var _ = Describe("yaml", func() {
	It("generates a map", func() {
		subMap := map[string]int{"c": 2}
		inputMap := pipeline.GenericMap{
			"a": 1,
			"b": subMap,
		}
		outputString, err := MapToString("yaml", inputMap)
		Expect(err).NotTo(HaveOccurred())
		Expect(outputString).To(Equal(`a: 1
b:
    c: 2
`))
	})
})
