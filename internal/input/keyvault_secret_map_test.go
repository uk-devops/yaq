package input_test

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/uk-devops/yaq/internal/input"
	"github.com/uk-devops/yaq/internal/testutils"
)

var _ = Describe("keyvault-secret-map", Label("azure"), func() {
	secretContent := "secret content"
	var kvName string
	const secretName = "yaq-keyvault-map-test"
	var client *azsecrets.Client

	BeforeEach(func() {
		kvName = testutils.KeyvaultName()
		client = testutils.KVClient(kvName)

		_, err := client.SetSecret(context.Background(), secretName, secretContent, nil)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		testutils.DeleteSecret(client, secretName)
	})

	Context("happy path", func() {
		It("returns data as a string", func() {
			inputString, err := input.ReadString("keyvault-secret-map:" + kvName + "/" + secretName)

			Expect(err).NotTo(HaveOccurred())
			Expect(inputString).To(Equal(secretContent))
		})
	})
})
