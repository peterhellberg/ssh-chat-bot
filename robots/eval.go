package robots

import (
	"fmt"
	"strings"
	"time"

	"github.com/dop251/goja"
)

// EvalBot is a bot that evaluates JavaScript
type EvalBot struct{}

func init() {
	RegisterRobot("eval", func() (robot Robot) { return new(EvalBot) })
}

// Run executes a command
func (b EvalBot) Run(c *Command) string {
	vm := goja.New()

	time.AfterFunc(1*time.Second, func() {
		vm.Interrupt("JS timeout")
	})

	value, err := vm.RunString(strings.Join(c.Args, " "))
	if err != nil {
		return err.Error()
	}

	out := fmt.Sprintf("%v", value)

	if strings.HasPrefix(out, "/") ||
		strings.Contains(out, "\r") ||
		strings.Contains(out, "\b") ||
		strings.Contains(out, "/me") ||
		strings.Contains(out, "/msg") ||
		strings.Contains(out, "/nick") ||
		strings.Contains(out, "exit") {
		return "Sorry, I won’t do that"
	}

	if len(out) > 128 {
		return "Output too big!"
	}

	return out
}

// Description describes what the robot does
func (b EvalBot) Description() string {
	return "<js>"
}
