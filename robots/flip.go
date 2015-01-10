package robots

import (
	"strings"

	"github.com/peterhellberg/flip"
)

// FlipBot is a bot who flips text
type FlipBot struct{}

func init() {
	RegisterRobot("flip", func() (robot Robot) { return new(FlipBot) })
}

// Run executes a deferred action
func (b FlipBot) Run(c *Command) string {
	switch c.Args[0] {
	default:
		return flip.UpsideDown(strings.Join(c.Args, " "))
	case "table", "t":
		return flip.Table(strings.Join(c.Args[1:], " "))
	case "gopher", "g":
		return flip.Gopher(strings.Join(c.Args[1:], " "))
	}

	return ""
}

// Description describes what the robot does
func (b FlipBot) Description() string {
	return "<table,gopher>"
}
