package robots

import "strings"

// EchoBot is a simple echo bot
type EchoBot struct{}

func init() {
	RegisterRobot("echo", func() (robot Robot) { return new(EchoBot) })
}

// Run executes a command
func (b EchoBot) Run(c *Command) string {
	return strings.Join(c.Args, " ")
}

// Description describes what the robot does
func (b EchoBot) Description() string {
	return "<something>"
}
