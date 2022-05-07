package output_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/output"
	"github.com/saliceti/yaq/pipeline"
)

var _ = Describe("Command", func() {
	var file *os.File

	Context("a valid command is passed", func() {
		BeforeEach(func() {
			var err error
			file, err = ioutil.TempFile(".", "data.*.yml")
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			os.Remove(file.Name())
		})

		It("runs the command with the environment variables", func() {
			err := PushMap("command", pipeline.GenericMap{"a": "1"}, []string{"bash", "-c", "echo -n $a > " + file.Name()})
			Expect(err).NotTo(HaveOccurred())

			data, err := ioutil.ReadFile(file.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal("1"))
		})
	})

	Context("no command is passed", func() {
		It("throws an error", func() {
			err := PushMap("command", pipeline.GenericMap{"a": "1"}, nil)
			Expect(err).To(MatchError("Empty command"))

		})
	})
})
