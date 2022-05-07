package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestYaq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Yaq Suite")
}
