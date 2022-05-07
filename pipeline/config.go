package pipeline

import (
	"bytes"
	"flag"
	"os"
	"strings"
)

// Config is a struct to store the initial configuration
type Config struct {
	Input  arrayFlags
	DumpTo string
	Output string
	Extra  []string
	Flags  flag.FlagSet
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join([]string(*i), ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var inputFlags arrayFlags

func GetConfig(progname string, args []string) (*Config, string, error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var outputBuffer bytes.Buffer
	flags.SetOutput(&outputBuffer)

	var config Config
	flags.Var(&config.Input, "i", "Pull from input (ex: stdin). -i may be repeated to fetch more inputs.")
	flags.StringVar(&config.DumpTo, "d", "json", "Dump to format (ex: json, yaml)")
	flags.StringVar(&config.Output, "o", "stdout", "Push to output (ex: stdout)")

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
