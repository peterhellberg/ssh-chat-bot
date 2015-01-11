package robots

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/robertkrimen/otto"
)

var halt = errors.New("Halt!")

// EvalBot is a bot that evaluates JavaScript
type EvalBot struct{}

func init() {
	RegisterRobot("eval", func() (robot Robot) { return new(EvalBot) })
}

// Run executes a command
func (b EvalBot) Run(c *Command) string {
	js := strings.Join(c.Args, " ")

	value, err := b.runUnsafe(js)
	if err != nil {
		return ""
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

	return out
}

// Description describes what the robot does
func (b EvalBot) Description() string {
	return "<js>"
}

func (b EvalBot) runUnsafe(unsafe string) (otto.Value, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		if caught := recover(); caught != nil {
			if caught == halt {
				fmt.Fprintf(os.Stderr, "Some code took to long! Stopping after: %v\n", duration)
				return
			}

			return
		}
		fmt.Fprintf(os.Stderr, "Ran code successfully: %v\n", duration)
	}()

	vm := otto.New()
	vm.Interrupt = make(chan func(), 1) // The buffer prevents blocking

	go func() {
		time.Sleep(2 * time.Second) // Stop after two seconds
		vm.Interrupt <- func() {
			panic(halt)
		}
	}()

	return vm.Run(unsafe) // Here be dragons (risky code)
}
