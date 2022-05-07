package testutils

import (
	"context"
	"fmt"
	"time"

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

func DeleteSecret(client *azsecrets.Client, secretName string) {
	resp, err := client.BeginDeleteSecret(context.TODO(), secretName, nil)
	Expect(err).NotTo(HaveOccurred())

	_, err = resp.PollUntilDone(context.TODO(), 250*time.Millisecond)
	Expect(err).NotTo(HaveOccurred())

	_, err = client.PurgeDeletedSecret(context.TODO(), secretName, nil)
	Expect(err).NotTo(HaveOccurred())
}
