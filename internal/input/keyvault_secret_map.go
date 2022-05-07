package input

import (
	"strings"

	"github.com/saliceti/yaq/internal/pipeline"
)

func init() {
	RegisterStringFunction("keyvault-secret-map", ReadSecretFromKeyVault)
}

func ReadSecretFromKeyVault(keyvaultSecret string) (string, error) {
	kvSecretArray := strings.Split(keyvaultSecret, "/")

	client, err := pipeline.KeyvaultClient(kvSecretArray[0])
	if err != nil {
		return "", err
	}

	secret, err := getSecret(client, kvSecretArray[1])
	if err != nil {
		return "", err
	}

	return secret, nil
}
