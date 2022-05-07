package output

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/saliceti/yaq/pipeline"
)

func init() {
	RegisterMapFunction("command", PushToCommandEnvVariables)
}

func PushToCommandEnvVariables(inputData interface{}, command []string) error {
	inputMap, ok := inputData.(pipeline.GenericMap)
	if !ok {
		return errors.New("command only accepts a map")
	}
	var cmd *exec.Cmd

	if len(command) == 0 {
		return errors.New("empty command")
	} else if len(command) == 1 {
		cmd = exec.Command(command[0])
	} else if len(command) > 1 {
		cmd = exec.Command(command[0], command[1:]...)
	}

	cmd.Env = commandEnvironment(inputMap)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err
}

func commandEnvironment(dataMap pipeline.GenericMap) []string {
	environment := make([]string, len(dataMap))
	i := 0

	for k, v := range dataMap {
		environment[i] = k + "=" + fmt.Sprintf("%v", v)
		i++
	}

	return environment
}
