package robots

// HelpBot is a robot that list all the robots
type HelpBot struct{}

func init() {
	RegisterRobot("help", func() (robot Robot) { return new(HelpBot) })
}

// Run returns a list of all robots
func (b HelpBot) Run(c *Command) string {
	output := ""

	for name, fn := range Robots {
		if name != "help" {
			output = output + name + "{" + fn().Description() + "} "
		}
	}

	return output
}

// Description describes what the robot does
func (b HelpBot) Description() string {
	return "List commands!"
}
