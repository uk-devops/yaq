package input

import (
	"io/ioutil"
	"os"

	"github.com/saliceti/yaq/internal/pipeline"
)

func init() {
	Register("stdin", ReadFromStdin)
}

func ReadFromStdin(_ string) (string, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		// stdin is empty and waiting for input from terminal
		// We don't allow typing input in the terminal
		return "", &pipeline.UsageError{Message: "Nothing to read from standard input"}
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
