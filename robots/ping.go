package robots

// PingBot is a simple ping/pong bot
type PingBot struct{}

func init() {
	RegisterRobot("ping", func() (robot Robot) { return new(PingBot) })
}

// Run executes a command
func (b PingBot) Run(c *Command) string {
	return "pong"
}

// Description describes what the robot does
func (b PingBot) Description() string {
	return ""
}
