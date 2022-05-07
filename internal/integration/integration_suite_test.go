package integration_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var binaryPath string

var _ = BeforeSuite(func() {
	binaryPath = build()
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func build() string {
	binaryPath, err := gexec.Build("github.com/saliceti/yaq")
	Expect(err).NotTo(HaveOccurred())

	return binaryPath
}

func runProgram(path string, args []string) *gexec.Session {
	cmd := exec.Command(path, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
