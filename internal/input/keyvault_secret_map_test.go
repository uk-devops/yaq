package input_test

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/saliceti/yaq/internal/input"
)

func kvClient(kvName string) *azsecrets.Client {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	Expect(err).NotTo(HaveOccurred())

	client, err := azsecrets.NewClient(
		fmt.Sprintf("https://%s.vault.azure.net/", kvName), cred, nil)
	Expect(err).NotTo(HaveOccurred())

	return client
}

var _ = Describe("keyvault-secret-map", Label("azure"), func() {
	secretContent := "secret content"
	var kvName string
	const secretName = "yaq-keyvault-map-test"

	BeforeEach(func() {
		var KEYVAULT_NAME_is_present bool
		kvName, KEYVAULT_NAME_is_present = os.LookupEnv("KEYVAULT_NAME")
		Expect(KEYVAULT_NAME_is_present).To(BeTrue())

		client := kvClient(kvName)

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
