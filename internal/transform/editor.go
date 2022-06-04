package transform

import (
	"io/ioutil"
	"os"
	"strings"

	"os/exec"

	"github.com/uk-devops/yaq/internal/pipeline"
	"gopkg.in/yaml.v3"
)

func init() {
	RegisterTransformFunction("editor", ProcessWithEditor)
}

func ProcessWithEditor(inputData pipeline.StructuredData, editor string) (pipeline.StructuredData, error) {
	yamlData, err := yaml.Marshal(inputData.Map())
	if err != nil {
		return nil, err
	}

	tempFile, err := ioutil.TempFile(".", "data.*.yml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	err = os.WriteFile(tempFile.Name(), yamlData, 0755)
	if err != nil {
		return nil, err
	}

	editorArray := strings.Split(editor, " ")

	var cmd *exec.Cmd
	if len(editorArray) == 1 {
		cmd = exec.Command(editorArray[0], tempFile.Name())
	} else {
		args := append(editorArray[1:], tempFile.Name())
		cmd = exec.Command(editorArray[0], args...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	newYamlData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return nil, err
	}
	var out pipeline.GenericMap
	err = yaml.Unmarshal(newYamlData, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
