package main

import (
	"flag"
	"log"
	"os"

	"github.com/uk-devops/yaq/internal/dump"
	"github.com/uk-devops/yaq/internal/input"
	"github.com/uk-devops/yaq/internal/load"
	"github.com/uk-devops/yaq/internal/output"
	"github.com/uk-devops/yaq/internal/pipeline"
	"github.com/uk-devops/yaq/internal/transform"
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

	var data, newData pipeline.StructuredData

	for _, inputArg := range config.Input {

		inputRequiresMap, err := input.CreatesMap(inputArg)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(10)
		}
		if inputRequiresMap {
			newData, err = input.ReadMap(inputArg)
			if err != nil {
				if _, ok := err.(*pipeline.UsageError); ok {
					log.Println(err)
					config.Flags.Usage()
					os.Exit(11)
				}
				log.Printf("Error: %v", err)
				os.Exit(12)
			}
		} else {
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

			newData, err = load.Unmarshal(inputString)
			if err != nil {
				log.Printf("Error: %v", err)
				os.Exit(4)
			}
		}

		if data == nil {
			data = newData
		} else {
			data = data.Append(newData)
		}
	}

	if config.Transform != "" {
		data, err = transform.TransformWith(config.Transform, data)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(8)
		}
	}

	outputRequiresMap, err := output.RequiresMap(config.Output)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(9)
	}

	if outputRequiresMap {
		err = output.PushMap(config.Output, data, config.Extra)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(7)
		}
	} else {
		outputString, err := dump.MapToString(config.DumpTo, data)
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
