package robots

import "strings"

// HelpBot is a robot that list all the robots
type HelpBot struct{}

func init() {
	RegisterRobot("help", func() (robot Robot) { return new(HelpBot) })
}

// Run returns a list of all robots
func (b HelpBot) Run(c *Command) string {
	outputs := []string{}

	for name, fn := range Robots {
		if name != "help" {
			output := name

			if desc := fn().Description(); desc != "" {
				output += " " + desc
			}

			outputs = append(outputs, output)
		}
	}

	return strings.Join(outputs, ", ")
}

// Description describes what the robot does
func (b HelpBot) Description() string {
	return ""
}
