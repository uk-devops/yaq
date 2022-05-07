package output_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/output"
)

const INPUT_DATA = "input data"

var _ = Describe("stdout", func() {
	Context("when a string is passed", func() {
		It("writes it to stdout", func() {
			r, w, err := os.Pipe()
			Expect(err).NotTo(HaveOccurred())

			origStdout := os.Stdout
			os.Stdout = w

			PushToStdout(INPUT_DATA)

			var output []byte
			output = make([]byte, len(INPUT_DATA))

			_, err = r.Read(output)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(output)).To(Equal(INPUT_DATA))

			os.Stdin = origStdout
		})
	})

})
