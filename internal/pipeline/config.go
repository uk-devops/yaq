package pipeline

import (
	"bytes"
	"flag"
	"os"
	"strings"
)

// Config is a struct to store the initial configuration
type Config struct {
	Input     arrayFlags
	Transform string
	DumpTo    string
	Output    string
	Extra     []string
	Flags     flag.FlagSet
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join([]string(*i), ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func GetConfig(progname string, args []string) (*Config, string, error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var outputBuffer bytes.Buffer
	flags.SetOutput(&outputBuffer)

	var config Config
	flags.Var(&config.Input, "i", `Pull from input. -i may be repeated to fetch more inputs. Available inputs:
-i stdin
-i file:/path/to/file
-i keyvault-secret-map:<keyvault-name>/<secret-name>
-i keyvault-secrets:<keyvault-name>
(required)`)
	flags.StringVar(&config.Transform, "t", "", `Apply transformation to data. Available transformations:
-t jq:"jq expression"
(default: No transformation)`)
	flags.StringVar(&config.DumpTo, "d", "json", `Dump to format. Available formats:
-d json
-d yaml
`)
	flags.StringVar(&config.Output, "o", "stdout", `Push to output. Available outputs:
-o stdout
-o file:/path/to/file
-o keyvault-secret:<keyvault-name>/<secret-name>
-o keyvault-secrets:<keyvault-name>
`)

	err := flags.Parse(args)
	if err != nil {
		return nil, outputBuffer.String(), err
	}

	config.Extra = flags.Args()
	config.Flags = *flags
	// Revert to stdout so Usage() can be called later on
	flags.SetOutput(os.Stdout)

	return &config, outputBuffer.String(), nil
}

func SplitArg(arg string) (string, string) {
	var name, parameter string
	argArray := strings.Split(arg, ":")
	name = argArray[0]
	if len(argArray) == 1 {
		parameter = ""
	} else {
		parameter = argArray[1]
	}
	return name, parameter
}
