package main

import (
	"encoding/json"
	"fmt"
	"github.com/apex/go-apex"
	"os"
	"os/exec"
)

type commandOutput struct {
  Cmd string `json:"cmd"`
	Output string `json:"output"`
  Error string `json:"error"`
  // exit code integer
  // error message
}

type commandInput struct {
	Cmd  string `json:"cmd"`
	Args string `json:"args"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var cmdIn commandInput
		json.Unmarshal(event, &cmdIn)
		fmt.Fprintf(os.Stderr, cmdIn.Cmd)
		fmt.Fprintf(os.Stderr, cmdIn.Args)

		var m commandOutput
		test := exec.Command(cmdIn.Cmd, cmdIn.Args)
		output, err := test.Output()
		m.Output = string(output)

    if (err != nil) {
      m.Error = err.Error()
    }

    m.Cmd = cmdIn.Cmd

		return m, nil
	})
}
