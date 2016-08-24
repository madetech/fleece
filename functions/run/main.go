package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
  "os/exec"
)

type commandOutput struct {
  Output string `json:"output"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
    var m commandOutput

    test := exec.Command("./test")
    output, _ := test.Output()
    m.Output = string(output)

		return m, nil
	})
}
