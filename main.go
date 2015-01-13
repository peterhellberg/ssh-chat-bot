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
)

var (
	user    = flag.String("n", "ssh-chat-bot", "Username")
	owner   = flag.String("o", "peterhellberg", "Bot owner username")
	host    = flag.String("h", "localhost", "Hostname")
	port    = flag.Int("p", 2200, "Port")
	verbose = flag.Bool("v", false, "Verbose output")
	delay   = flag.Duration("d", 5*time.Second, "Delay")
)

const repoURL = "https://github.com/peterhellberg/ssh-chat-bot"

var buildCommit string

func main() {
	flag.Usage = usage
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	if err := Bot(addr); err != nil {
		l("Error: %v", err)
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

func l(format string, args ...interface{}) {
	if *verbose {
		fmt.Printf(format+"\n", args...)
	}
}
