package output

import (
	"fmt"
)

func init() {
	RegisterStringFunction("stdout", PushToStdout)
}

func PushToStdout(inputString, _ string) error {
	fmt.Println(inputString)
	return nil
}
