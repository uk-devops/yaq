package input

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func init() {
	RegisterStringFunction("keyvault-secret-map", ReadFromKeyVault)
}

func ReadFromKeyVault(keyvaultSecret string) (string, error) {
	kvSecretArray := strings.Split(keyvaultSecret, "/")

	client, err := kvClient(kvSecretArray[0])
	if err != nil {
		return "", err
	}

	secret, err := getSecret(client, kvSecretArray[1])
	if err != nil {
		return "", err
	}

	return secret, nil
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

func getSecret(client *azsecrets.Client, secretName string) (string, error) {
	response, err := client.GetSecret(context.Background(), secretName, nil)
	if err != nil {
		return "", err
	}

	return *response.Value, nil
}
