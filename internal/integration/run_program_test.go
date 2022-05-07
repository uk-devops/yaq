package integration_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Build and run", func() {
	var session *gexec.Session

	Context("no argument", func() {
		BeforeEach(func() {
			session = runProgram(binaryPath, []string{})
		})

		It("exits with status 0", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints usage", func() {
			Eventually(session).Should(gbytes.Say("Usage"))
		})
	})

	Context("error in arguments", func() {
		BeforeEach(func() {
			session = runProgram(binaryPath, []string{"-i"})
		})

		It("exits with status 1", func() {
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints usage", func() {
			Eventually(session.Err).Should(gbytes.Say("flag needs an argument: -i"))
		})
	})

	Context("1 input", func() {
		var file *os.File
		fileData := []byte(`brand: peugeot`)

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

		It("prints the data as json", func() {
			Eventually(session.Out).Should(gbytes.Say(`{
  "brand": "peugeot"
}`))
		})
	})

	Context("2 input", func() {
		var file1 *os.File
		var file2 *os.File
		file1Data := []byte(`brand: peugeot`)
		file2Data := []byte(`model: 205`)

		BeforeEach(func() {
			var err error
			file1, err = ioutil.TempFile(".", "data.*.yml")
			Expect(err).NotTo(HaveOccurred())
			file2, err = ioutil.TempFile(".", "data.*.yml")
			Expect(err).NotTo(HaveOccurred())

			err = os.WriteFile(file1.Name(), file1Data, 0755)
			Expect(err).NotTo(HaveOccurred())
			err = os.WriteFile(file2.Name(), file2Data, 0755)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			os.Remove(file1.Name())
			os.Remove(file2.Name())
		})
		BeforeEach(func() {
			session = runProgram(binaryPath, []string{"-i", "file:" + file1.Name(), "-i", "file:" + file2.Name()})
		})

		It("exits with status 0", func() {
			Eventually(session).Should(gexec.Exit(0))
		})

		It("prints the merged data as json", func() {
			Eventually(session.Out).Should(gbytes.Say(`{
  "brand": "peugeot",
  "model": 205
}`))
		})
	})
})
