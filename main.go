package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"os"
	"os/exec"
)

func main() {
	j := []byte(`{"yaml": "test", "yaml2": test2}`)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "YAML_OUTPUT", "--value", string(y)).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}

	os.Exit(0)
}
