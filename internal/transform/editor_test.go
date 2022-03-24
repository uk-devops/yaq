package transform_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/uk-devops/yaq/internal/pipeline"
	"github.com/uk-devops/yaq/internal/transform"
)

var _ = Describe("editor", func() {
	var output pipeline.StructuredData
	var err error
	mapInput := pipeline.GenericMap{"a": 1, "b": "c"}

	Context("map input", func() {
		var tempEditor *os.File

		BeforeEach(func() {
			// Create fake editor as mini bash script
			tempEditor, err = ioutil.TempFile(".", "editor_*.sh")
			Expect(err).NotTo(HaveOccurred())

			editorString := `awk '{gsub("c","d");print $0}' $1 > new
mv new $1
`
			err = os.WriteFile(tempEditor.Name(), []byte(editorString), 0755)
			Expect(err).NotTo(HaveOccurred())
		})

		BeforeEach(func() {
			args := "bash ./" + tempEditor.Name()
			output, err = transform.TransformWith("editor:"+args, mapInput)
			Expect(err).NotTo(HaveOccurred())
		})

		It("transforms the map", func() {
			Expect(output).To(Equal(pipeline.GenericMap{"a": 1, "b": "d"}))
		})

		AfterEach(func() {
			os.Remove(tempEditor.Name())
		})
	})
})
