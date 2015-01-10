package robots

// AboutBot describes ssh-chat-bot
type AboutBot struct{}

func init() {
	RegisterRobot("about", func() (robot Robot) { return new(AboutBot) })
}

// Run executes a command
func (b AboutBot) Run(c *Command) string {
	return "https://github.com/peterhellberg/ssh-chat-bot"
}

// Description describes what the robot does
func (b AboutBot) Description() string {
	return ""
}
