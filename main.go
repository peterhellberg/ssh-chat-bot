/*

ssh-chat-bot

A small chatbot for ssh-chat

*/
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/peterhellberg/ssh-chat-bot/herder"
)

const repoURL = "https://github.com/peterhellberg/ssh-chat-bot"

var buildCommit string

func main() {
	var (
		user    = flag.String("n", "ssh-chat-bot", "Username")
		owner   = flag.String("o", "peter", "Bot owner username")
		host    = flag.String("h", "localhost", "Hostname")
		port    = flag.Int("p", 2022, "Port")
		delay   = flag.Duration("d", 2*time.Second, "Delay")
		check   = flag.Duration("c", 30*time.Second, "Duration between alive checks")
		verbose = flag.Bool("v", false, "Verbose output")
	)

	flag.Usage = usage

	flag.Parse()

	h := herder.New(
		herder.User(*user),
		herder.Owner(*owner),
		herder.Addr(fmt.Sprintf("%s:%d", *host, *port)),
		herder.Delay(*delay),
		herder.Check(*check),
		herder.Verbose(*verbose),
	)

	if err := h.Run(); err != nil {
		if *verbose {
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ./ssh-chat-bot [-h hostname] [-v]\n\n")

	if buildCommit != "" {
		fmt.Fprintf(os.Stderr, "build: "+repoURL+"/commit/"+buildCommit+"\n\n")
	}

	fmt.Fprintf(os.Stderr, "flags:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(2)
}
