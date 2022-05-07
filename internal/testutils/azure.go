package testutils

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	. "github.com/onsi/gomega"
)

func KVClient(kvName string) *azsecrets.Client {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	Expect(err).NotTo(HaveOccurred())

	client, err := azsecrets.NewClient(
		fmt.Sprintf("https://%s.vault.azure.net/", kvName), cred, nil)
	Expect(err).NotTo(HaveOccurred())

	return client
}
