package load_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLoad(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Load Suite")
}
