package input

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func getSecret(client *azsecrets.Client, secretName string) (string, error) {
	response, err := client.GetSecret(context.Background(), secretName, nil)
	if err != nil {
		return "", err
	}

	return *response.Value, nil
}
