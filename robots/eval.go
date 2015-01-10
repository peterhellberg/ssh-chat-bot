package robots

import (
	"fmt"
	"strings"

	"github.com/robertkrimen/otto"
)

// EvalBot is a bot that evaluates JavaScript
type EvalBot struct {
	// VM is an Otto VM
	VM *otto.Otto
}

func init() {
	RegisterRobot("eval", func() (robot Robot) {
		return &EvalBot{VM: otto.New()}
	})
}

// Run executes a command
func (b EvalBot) Run(c *Command) string {
	js := strings.Join(c.Args, " ")

	value, err := b.VM.Run(js)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}

// Description describes what the robot does
func (b EvalBot) Description() string {
	return "<js>"
}
