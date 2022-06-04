# yaq: yet another jq, yq...

Utility to collect structured data and push it somewhere else it any format.

For example, convert from json to yaml:

```shell
% echo '{"code": 1234}' | yaq -d yaml
code: 1234
```

This command will:
- Pull data from standard input
- Load json data
- Convert to yaml data
- Push to standard out

It can also...

### Read from yaml or json
Convert from yaml to json:

```shell
% echo 'code: 1234' | yaq -d json
{
  "code": 1234
}
```

### Read from standard input
```shell
% echo '{"code": 1234}' | yaq -i stdin
```

### Read from file
```shell
% yaq -i file:input.yml
```

### Convert to yaml
```shell
% yaq -i file:input.json -d yaml
```

### Convert to json (default)
```shell
% yaq -i file:input.yml -d json
```

### Read from multiple sources
```shell
% yaq -i file:input.yml -i file:input.json
```
If some keys are the same, the last one always wins.

### Write to file
```shell
% yaq -i file:input.yml -i file:input.json -d yaml -o file:output.yml
```

### Run a command and populate its environment
```shell
% yaq -i file:input.yml -o command -- bash
bash-3.2$ echo $code
1234
```

### Apply a transformation with jq syntax
```shell
% curl -s https://api.chucknorris.io/jokes/random  | yaq -i stdin -t jq:'.value | ascii_upcase' -o command -- bash -c 'echo $result && say $result'
```

### Read a map from an Azure keyvault secret
```shell
% yaq -i keyvault-secret-map:myKeyVault/MySecret -d yaml
```

### Write a map to an Azure keyvault secret
```shell
% yaq -i file:input.yml -i file:input.json -o keyvault-secret:myKeyvault/MySecret
```

### Read and write all secrets in an Azure keyvault
```shell
% yaq -i keyvault-secrets:myKeyvault -d yaml -o file:input.yml
```

```shell
% yaq -i file:input.yml -o keyvault-secrets:myKeyvault
```

### Edit secrets interactively
```shell
% yaq -i keyvault-secret-map:myKeyvault/mySecret -d yaml -t editor:vim -o keyvault-secret:myKeyvault/mySecret
<Opens vim to edit the data as yaml. And saves it back to the Keyvault secret if the syntax is correct>
```

## Install

Prerequisites: [go](https://go.dev) environment

run: `make install`

## Run unit tests

run: `make tests`

## Run the Azure tests

- Create an Azure [key vault](https://azure.microsoft.com/en-gb/services/key-vault/#product-overview)
- Configure the access policy or RBAC so the user has full access to secrets (including purge)
- Login via az cli
- Export the key vault name and run the Azure tests

  ```
  % KEYVAULT_NAME=myYaqTestKeyVault make azure-test
  ```
