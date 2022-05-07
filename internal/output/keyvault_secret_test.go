package output_test

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/saliceti/yaq/internal/output"
	"github.com/saliceti/yaq/internal/testutils"
)

var _ = Describe("keyvault-secret-map", Label("azure"), func() {

	Context("happy path", func() {
		secretContent := "secret content"
		var kvName string
		const secretName = "yaq-keyvault-map-test"
		var client *azsecrets.Client

		BeforeEach(func() {
			var KEYVAULT_NAME_is_present bool
			kvName, KEYVAULT_NAME_is_present = os.LookupEnv("KEYVAULT_NAME")
			Expect(KEYVAULT_NAME_is_present).To(BeTrue())

			err := output.PushString("keyvault-secret:"+kvName+"/"+secretName, secretContent)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			testutils.DeleteSecret(client, secretName)
		})

		It("writes a string into the secret", func() {
			client = testutils.KVClient(kvName)

			response, err := client.GetSecret(context.Background(), secretName, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(*response.Value).To(Equal(secretContent))
		})
	})
})
