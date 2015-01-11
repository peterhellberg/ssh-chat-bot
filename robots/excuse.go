package robots

import "github.com/PuerkitoBio/goquery"

// ExcuseBot is a simple echo bot
type ExcuseBot struct{}

func init() {
	RegisterRobot("excuse", func() (robot Robot) { return new(ExcuseBot) })
}

// Run executes a command
func (b ExcuseBot) Run(c *Command) string {
	doc, err := goquery.NewDocument("http://developerexcuses.com/")
	if err != nil {
		return ""
	}

	return doc.Find(".wrapper center a").First().Text()
}

// Description describes what the robot does
func (b ExcuseBot) Description() string {
	return ""
}
