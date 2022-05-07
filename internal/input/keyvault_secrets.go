package input

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/saliceti/yaq/internal/pipeline"
)

func init() {
	RegisterMapFunction("keyvault-secrets", ReadAllSecretsFromKeyVault)
}

func ReadAllSecretsFromKeyVault(keyvaultName string) (pipeline.StructuredData, error) {
	client, err := pipeline.KeyvaultClient(keyvaultName)
	if err != nil {
		return nil, err
	}

	secretMap := pipeline.GenericMap{}

	for _, name := range listSecrets(client) {
		secret, err := getSecret(client, name)
		if err != nil {
			return nil, err
		}
		secretMap[name] = secret
	}

	return secretMap, nil
}

func listSecrets(client *azsecrets.Client) []string {
	var secretNames []string
	pager := client.ListSecrets(nil)
	var idArray []string
	var secretName string

	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		for _, v := range page.Secrets {
			idArray = strings.Split(*v.ID, "/")
			secretName = idArray[len(idArray)-1]
			secretNames = append(secretNames, secretName)
		}
	}

	return secretNames
}
