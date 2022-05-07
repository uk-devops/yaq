package input_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/saliceti/yaq/internal/input"
)

const INPUT_DATA = "input data"

var _ = Describe("stdin", func() {
	Context("it receives data", func() {
		It("returns data as a string", func() {
			r, w, err := os.Pipe()
			Expect(err).NotTo(HaveOccurred())

			origStdin := os.Stdin
			os.Stdin = r

			w.Write([]byte(INPUT_DATA))
			w.Close()

			input_string, err := ReadString("stdin")

			Expect(input_string).To(Equal(INPUT_DATA))
			Expect(err).NotTo(HaveOccurred())

			os.Stdin = origStdin
		})
	})
	Context("no data is sent", func() {
		It("throws an error", func() {
			input_string, err := ReadString("stdin")

			Expect(input_string).To(Equal(""))
			Expect(err).To(MatchError("Nothing to read from standard input"))
		})
	})

})
