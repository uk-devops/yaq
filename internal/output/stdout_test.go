package output_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/internal/output"
)

var _ = Describe("stdout", func() {
	const INPUT_DATA = "input data"

	Context("when a string is passed", func() {
		It("writes it to stdout", func() {
			r, w, err := os.Pipe()
			Expect(err).NotTo(HaveOccurred())

			origStdout := os.Stdout
			os.Stdout = w

			err = PushString("stdout", INPUT_DATA)
			Expect(err).NotTo(HaveOccurred())

			output := make([]byte, len(INPUT_DATA))

			_, err = r.Read(output)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(output)).To(Equal(INPUT_DATA))

			os.Stdin = origStdout
		})
	})

})
