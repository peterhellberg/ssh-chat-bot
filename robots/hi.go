package robots

import "github.com/peterhellberg/hi"

// HiBot is a bot that finds images for a given hashtag
type HiBot struct{}

func init() {
	RegisterRobot("hi", func() (robot Robot) { return new(HiBot) })
}

// Run executes a deferred action
func (b HiBot) Run(c *Command) string {
	if len(c.Args) > 0 {
		img, err := hi.FindShuffledImage(c.Args[0])
		if err == nil {
			return img.URL
		}
	}

	return ""
}

// Description describes what the robot does
func (b HiBot) Description() string {
	return "<hashtag>"
}
