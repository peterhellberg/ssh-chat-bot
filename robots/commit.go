package robots

import (
	"io/ioutil"
	"net/http"
)

// CommitBot replies with fun commit messages
type CommitBot struct{}

func init() {
	RegisterRobot("commit", func() (robot Robot) { return new(CommitBot) })
}

// Run executes a command
func (b CommitBot) Run(c *Command) string {
	resp, err := http.Get("http://whatthecommit.com/index.txt")
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

// Description describes what the robot does
func (b CommitBot) Description() string {
	return ""
}
