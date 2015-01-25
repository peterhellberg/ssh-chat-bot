# ssh-chat-bot

A small chatbot for [ssh-chat](https://github.com/shazow/ssh-chat).

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/ssh-chat-bot)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/ssh-chat-bot#mit-license)

## Installation

```bash
go get -u github.com/peterhellberg/ssh-chat-bot
```

You can also clone the repo and then run `make` in
order to populate the `main.buildCommit` variable.

## Usage

```bash
usage: ./ssh-chat-bot [-h hostname] [-v]

flags:
  -c=30s: Duration between alive checks
  -d=5s: Delay
  -h="localhost": Hostname
  -n="ssh-chat-bot": Username
  -o="peterhellberg": Bot owner username
  -p=2200: Port
  -v=false: Verbose output
```

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
