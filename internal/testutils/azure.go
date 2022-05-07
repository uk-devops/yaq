package testutils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"

	//lint:ignore ST1001 this is only used for tests
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

func DeleteSecret(client *azsecrets.Client, secretName string) {
	resp, err := client.BeginDeleteSecret(context.TODO(), secretName, nil)
	Expect(err).NotTo(HaveOccurred())

	_, err = resp.PollUntilDone(context.TODO(), 250*time.Millisecond)
	Expect(err).NotTo(HaveOccurred())

	_, err = client.PurgeDeletedSecret(context.TODO(), secretName, nil)
	Expect(err).NotTo(HaveOccurred())
}

func KeyvaultName() string {
	kvName, KEYVAULT_NAME_is_present := os.LookupEnv("KEYVAULT_NAME")
	Expect(KEYVAULT_NAME_is_present).To(BeTrue())
	return kvName
}
