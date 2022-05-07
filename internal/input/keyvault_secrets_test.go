package input_test

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/uk-devops/yaq/internal/input"
	"github.com/uk-devops/yaq/internal/pipeline"
	"github.com/uk-devops/yaq/internal/testutils"
)

var _ = Describe("keyvault-secrets", Label("azure"), func() {
	var kvName string
	secretNames := []string{"yaq-keyvault-map-test-0", "yaq-keyvault-map-test-1"}
	secretContent := []string{"secret-content-0", "secret-content-1"}
	var client *azsecrets.Client

	BeforeEach(func() {
		kvName = testutils.KeyvaultName()
		client = testutils.KVClient(kvName)

		for i, s := range secretNames {
			_, err := client.SetSecret(context.Background(), s, secretContent[i], nil)
			Expect(err).NotTo(HaveOccurred())
		}
	})

	AfterEach(func() {
		for _, s := range secretNames {
			testutils.DeleteSecret(client, s)
		}
	})

	Context("happy path", func() {
		It("returns a map of all secrets", func() {
			inputMap, err := input.ReadMap("keyvault-secrets:" + kvName)

			Expect(err).NotTo(HaveOccurred())

			expectedMap := pipeline.GenericMap{
				"yaq-keyvault-map-test-0": "secret-content-0",
				"yaq-keyvault-map-test-1": "secret-content-1",
			}
			Expect(inputMap).To(Equal(expectedMap))
		})
	})
})
