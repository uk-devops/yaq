package output

import (
	"fmt"
)

func init() {
	RegisterStringFunction("stdout", PushToStdout)
}

func PushToStdout(inputString string) {
	fmt.Println(inputString)
}
