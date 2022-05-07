package input_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/saliceti/yaq/internal/input"
)

var _ = Describe("File", func() {
	var file *os.File
	fileData := []byte(`brand: peugeot
model: 205
colour: blue
engine:
cylinders: 4
power: 300`)

	BeforeEach(func() {

		var err error
		file, err = ioutil.TempFile(".", "data.*.yml")
		Expect(err).NotTo(HaveOccurred())

		err = os.WriteFile(file.Name(), fileData, 0755)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.Remove(file.Name())
	})

	Context("the file can be opened", func() {
		It("returns data as a string", func() {
			inputString, err := ReadString("file:" + file.Name())

			Expect(err).NotTo(HaveOccurred())
			Expect(inputString).To(Equal(string(fileData)))
		})
	})

	Context("the file cannot be opened", func() {
		It("throws an error", func() {
			inputString, err := ReadString("file:unknown")

			Expect(err.Error()).To(Equal("open unknown: no such file or directory"))
			Expect(inputString).To(Equal(""))
		})
	})
})
