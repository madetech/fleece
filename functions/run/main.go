package main

import (
	"encoding/json"
	"fmt"
	"github.com/apex/go-apex"
	"os"
	"os/exec"
)

type commandOutput struct {
	Output string `json:"output"`
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
		test := exec.Command(fmt.Sprintf("./%s", cmdIn.Cmd))
		output, _ := test.Output()
		m.Output = string(output)

		return m, nil
	})
}
