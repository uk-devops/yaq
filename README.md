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

## Download
Select the package for your system from the [latest release](https://github.com/uk-devops/yaq/releases/latest). Unzip and add `yaq` to your PATH.


## It can also...

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

```shell
% yaq -i file:input.yml -o command -- ruby -e 'p ENV'
{"code"=>1234}
```

### Consume an API as environment variables in any process
Open latest xkcd [Mac]
```shell
% curl -s "https://xkcd.com/info.0.json" | yaq -o command -- bash -c 'open $img'
```

Display your IP
```shell
% curl -s "https://api.ipify.org/?format=json" | yaq -o command -- ruby -e 'puts "Your ip is: " + ENV["ip"]'
```

Say Chuck Norris fact [Mac]
```shell
% curl -s https://api.chucknorris.io/jokes/random  | yaq -o command -- bash -c 'say $value'
```

### Apply a transformation with jq syntax
List of bank holidays in UK in 2022
```shell
% curl -s https://date.nager.at/api/v2/publicholidays/2022/GB | yaq -t jq:'.[].date' -d yaml
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

### Copy all secrets between Azure keyvaults
```shell
% yaq -i keyvault-secrets:myKeyvault1 -o keyvault-secrets:myKeyvault2
```

### Edit secrets interactively
```shell
% yaq -i keyvault-secret-map:myKeyvault/mySecret -d yaml -t editor:vim -o keyvault-secret:myKeyvault/mySecret
<Opens vim to edit the data as yaml. And saves it back to the Keyvault secret if the syntax is correct>
```

## Install from source

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
