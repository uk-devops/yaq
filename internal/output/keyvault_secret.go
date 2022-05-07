package output

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func init() {
	RegisterStringFunction("keyvault-secret", WriteToKeyVault)
}

func WriteToKeyVault(outputString string, keyvaultSecret string) error {
	kvSecretArray := strings.Split(keyvaultSecret, "/")

	client, err := kvClient(kvSecretArray[0])
	if err != nil {
		return err
	}

	err = setSecret(client, kvSecretArray[1], outputString)
	if err != nil {
		return err
	}

	return nil
}

func kvClient(kvName string) (*azsecrets.Client, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	client, err := azsecrets.NewClient(
		fmt.Sprintf("https://%s.vault.azure.net/", kvName), cred, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func setSecret(client *azsecrets.Client, secretName string, value string) error {
	_, err := client.SetSecret(context.Background(), secretName, value, nil)
	if err != nil {
		return err
	}

	return nil
}
