package output

import (
	"context"
	"strings"

	"github.com/saliceti/yaq/internal/pipeline"
)

func init() {
	RegisterStringFunction("keyvault-secret", WriteToKeyVault)
}

func WriteToKeyVault(outputString string, keyvaultSecret string) error {
	kvSecretArray := strings.Split(keyvaultSecret, "/")

	client, err := pipeline.KeyvaultClient(kvSecretArray[0])
	if err != nil {
		return err
	}

	_, err = client.SetSecret(context.Background(), kvSecretArray[1], outputString, nil)
	if err != nil {
		return err
	}

	return nil
}
