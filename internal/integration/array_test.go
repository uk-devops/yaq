package integration_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Array", func() {
	var session *gexec.Session

	Context("1 array input", func() {
		var file *os.File
		fileData := []byte(`- brand: peugeot
- brand: renault`)

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
		BeforeEach(func() {
			session = runProgram(binaryPath, []string{"-i", "file:" + file.Name()})
		})

		It("exits with status 0", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints usage", func() {
			Eventually(session.Out).Should(gbytes.Say(`[
  {"brand": "peugeot"},
  {"brand": "renault"}
]`))
		})
	})
})
