package robots

import (
	"strings"

	"github.com/peterhellberg/giphy"
)

// GiphyBot is a robot who speak in GIFs
type GiphyBot struct {
	// Client is a client for the Giphy API
	Client *giphy.Client
}

func init() {
	RegisterRobot("giphy", func() (robot Robot) {
		return &GiphyBot{Client: giphy.DefaultClient}
	})
}

// Run executes a deferred action
func (b GiphyBot) Run(c *Command) string {
	args := []string{}

	for _, a := range c.Args {
		args = append(args, strings.ToLower(a))
	}

	switch args[0] {
	default:
		return b.search(c, args)
	case "search", "s":
		return b.search(c, args[1:])
	case "gif", "id":
		return b.gif(c, args[1:])
	case "random", "rand", "r":
		return b.random(c, args[1:])
	case "translate", "trans", "t":
		return b.translate(c, args[1:])
	case "trending", "trend", "tr":
		return b.trending(c, args[1:])
	}

	return ""
}

// Description describes what the robot does
func (b GiphyBot) Description() string {
	return "gifs!"
}

func (b GiphyBot) search(c *Command, args []string) string {
	res, err := b.Client.Search(args)
	if err != nil || len(res.Data) == 0 {
		return ""
	}

	return strings.Join(args, " ") + ": " + res.Data[0].Images.FixedHeight.URL
}

func (b GiphyBot) gif(c *Command, args []string) string {
	if len(args) == 0 {
		return ""
	}

	res, err := b.Client.GIF(args[0])
	if err != nil {
		return err.Error()
	}

	return "GIF: " + res.Data.Images.FixedHeightDownsampled.URL
}

func (b GiphyBot) random(c *Command, args []string) string {
	res, err := b.Client.Random(args)
	if err != nil {
		return err.Error()
	}

	return "Random: " + res.Data.FixedHeightDownsampledURL
}

func (b GiphyBot) translate(c *Command, args []string) string {
	res, err := b.Client.Translate(args)
	if err != nil {
		return err.Error()
	}

	return res.Data.Images.FixedHeight.URL
}

func (b GiphyBot) trending(c *Command, args []string) string {
	res, err := b.Client.Trending(args)
	if err != nil {
		return err.Error()
	}

	return "Trending: " + res.Data[0].Images.FixedHeight.URL
}
