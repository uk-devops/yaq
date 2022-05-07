package input_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/saliceti/yaq/internal/input"
	"github.com/saliceti/yaq/internal/testutils"
)

var _ = Describe("keyvault-secret-map", Label("azure"), func() {
	secretContent := "secret content"
	var kvName string
	const secretName = "yaq-keyvault-map-test"

	BeforeEach(func() {
		var KEYVAULT_NAME_is_present bool
		kvName, KEYVAULT_NAME_is_present = os.LookupEnv("KEYVAULT_NAME")
		Expect(KEYVAULT_NAME_is_present).To(BeTrue())

		client := testutils.KVClient(kvName)

		_, err := client.SetSecret(context.Background(), secretName, secretContent, nil)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("happy path", func() {
		It("returns data as a string", func() {
			inputString, err := input.ReadString("keyvault-secret-map:" + kvName + "/" + secretName)

			Expect(err).NotTo(HaveOccurred())
			Expect(inputString).To(Equal(secretContent))
		})
	})
})
