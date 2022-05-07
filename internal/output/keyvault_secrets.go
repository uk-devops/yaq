package output

import (
	"context"
	"fmt"

	"github.com/saliceti/yaq/internal/pipeline"
)

func init() {
	RegisterMapFunction("keyvault-secrets", PushAllSecretsToKeyVault)
}

func PushAllSecretsToKeyVault(secretData pipeline.StructuredData, keyvaultName string, _ []string) error {
	client, err := pipeline.KeyvaultClient(keyvaultName)
	if err != nil {
		return err
	}

	for s, v := range secretData.Map() {
		_, err := client.SetSecret(context.Background(), s, fmt.Sprintf("%v", v), nil)

		if err != nil {
			return err
		}
	}

	return nil
}
