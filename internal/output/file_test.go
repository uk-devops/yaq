package output_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/uk-devops/yaq/internal/output"
)

const INPUT_DATA = "input data"

var _ = Describe("file", func() {
	Context("when a string is passed", func() {
		Context("and a path is passed", func() {
			Context("and the path is writable", func() {
				var file *os.File
				var err error

				BeforeEach(func() {
					file, err = ioutil.TempFile(".", "data.*.yml")
					Expect(err).NotTo(HaveOccurred())
				})

				AfterEach(func() {
					os.Remove(file.Name())
				})

				BeforeEach(func() {
					PushString("file:"+file.Name(), INPUT_DATA)
				})

				It("writes it to a file", func() {
					data, err := ioutil.ReadFile(file.Name())
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(Equal(INPUT_DATA))
				})
			})

			Context("the path is not writable", func() {
				var err error

				BeforeEach(func() {
					err = PushString("file:/non/existent/path/to/file.txt", INPUT_DATA)
				})

				It("returns an error", func() {
					Expect(err).To(MatchError("can't write to file: open /non/existent/path/to/file.txt: no such file or directory"))
				})
			})
		})

		Context("and no path is passed", func() {
			It("returns an error", func() {
				err := PushString("file", INPUT_DATA)
				Expect(err).To(MatchError("output file path cannot be empty"))
			})
		})
	})
})
