# ssh-chat-bot

A simple chat-bot Used to check if a [ssh-chat](https://github.com/shazow/ssh-chat)
server is up, and responding.

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/ssh-chat-bot)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/ssh-chat-bot#mit-license)

## Robots

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/ssh-chat-bot/robots)

### Robot skeleton

```go
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
	return "pong"
}

```

## Usage

```bash
Usage of ./ssh-chat-bot:
  -n="ssh-chat-bot": Username
  -h="localhost": Hostname
  -p=2200: Port
  -v=false: Verbose output
```

## MIT License

*Copyright (C) 2015 [Peter Hellberg](http://c7.se/)*

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
