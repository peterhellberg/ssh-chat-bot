package robots

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// QuoteBot replies with a quote
type QuoteBot struct{}

func init() {
	RegisterRobot("quote", func() (robot Robot) { return new(QuoteBot) })
}

// Run executes a command
func (b QuoteBot) Run(c *Command) string {
	baseURL := "http://www.iheartquotes.com/api/v1/random?"
	baseParams := "max_lines=1&max_characters=256&show_permalink=false&source="
	source := ""

	if len(c.Args) > 0 {
		switch c.Args[0] {
		case "futurama":
			source = "futurama"
		case "fortune":
			source = "fortune"
		case "codehappy":
			source = "codehappy"
		}
	}

	resp, err := http.Get(baseURL + baseParams + source)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return strings.Replace(strings.Replace(strings.Replace(
		string(body), "\r", "", -1),
		"[", " [", -1),
		"&quot;", `"`, -1)
}

// Description describes what the robot does
func (b QuoteBot) Description() string {
	return "<futurama|fortune|codehappy>"
}
