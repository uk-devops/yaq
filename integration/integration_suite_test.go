package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
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
