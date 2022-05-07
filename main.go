package main

import (
	"flag"
	"log"
	"os"

	"github.com/saliceti/yaq/dump"
	"github.com/saliceti/yaq/input"
	"github.com/saliceti/yaq/load"
	"github.com/saliceti/yaq/output"
	"github.com/saliceti/yaq/pipeline"
)

func main() {
	config, cliOutput, err := pipeline.GetConfig(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(cliOutput)
		os.Exit(2)
	} else if err != nil {
		log.Println("Error:")
		log.Println(cliOutput)
		os.Exit(1)
	}

	if config.Input == nil {
		config.Flags.Usage()
		os.Exit(0)
	}

	data := pipeline.StructuredData{}

	for _, inputArg := range config.Input {

		inputString, err := input.ReadString(inputArg)
		if err != nil {
			if _, ok := err.(*pipeline.UsageError); ok {
				log.Println(err)
				config.Flags.Usage()
				os.Exit(7)
			}
			log.Printf("Error: %v", err)
			os.Exit(3)
		}

		newData, err := load.Unmarshal(inputString)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(4)
		}

		data.Append(newData)
	}

	if output.RequiresMap(config.Output) {
		err = output.PushMap(config.Output, data.Data, config.Extra)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(7)
		}
	} else {
		outputString, err := dump.MapToString(config.DumpTo, data.Data)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(5)
		}

		err = output.PushString(config.Output, outputString)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(6)
		}
	}
}
