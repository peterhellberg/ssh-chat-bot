package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/peterhellberg/ssh-chat-bot/robots"

	"golang.org/x/crypto/ssh"
)

const controlCodeString = "\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[D\x1b[K"

var active = false

// Bot runs the bot
func Bot(addr string) error {
	conn, err := dial(addr, *user)
	if err != nil {
		return err
	}

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	in, err := session.StdinPipe()
	if err != nil {
		return err
	}

	out, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	if err := session.Shell(); err != nil {
		return err
	}

	err = session.RequestPty("xterm", 80, 40, ssh.TerminalModes{})
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(*delay)
		in.Write([]byte("Now active\r\n"))
		active = true
	}()

	in.Write([]byte("/theme mono\r\n"))

	go func() {
		for {
			in.Write([]byte("/motd\r\n"))
			time.Sleep(30 * time.Second)
		}
	}()

	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return err
		}

		if strings.Contains(line, " "+*user+": ") {
			cmd, err := parseLine(line)
			if err == nil {
				robot, err := getRobot(cmd.Command)
				if err != nil {
					continue
				}

				if !active {
					continue
				}

				if response := robot.Run(cmd); response != "" {
					reply(in, fmt.Sprintf("%s %s", cmd.From, response))
				}
			}
		}
	}

	return errors.New("ERROR")
}

func parseLine(line string) (*robots.Command, error) {
	fields := strings.Fields(line)

	if len(fields) < 4 {
		return nil, errors.New("not enough fields in line")
	}

	fromFields := strings.Split(fields[1], controlCodeString)
	if len(fromFields) < 2 {
		return nil, errors.New("not enough fields in line")
	}
	from := fromFields[1]

	if len(fields) < 4 {
		return nil, errors.New("not enough fields in line")
	}

	command := fields[3]
	args := []string{}
	if len(fields) > 4 {
		args = fields[4:]
	}

	if active {
		fmt.Printf("%#v\n", args)
	}

	cmd := robots.Command{
		From:    from,
		Command: command,
		Args:    args,
	}

	return &cmd, nil
}

func reply(in io.WriteCloser, s string) {
	in.Write([]byte(s + "\r\n"))
}

func dial(addr, user string) (*ssh.Client, error) {
	key, err := MakeKey()
	if err != nil {
		return nil, err
	}

	return ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	})
}

func getRobot(command string) (robots.Robot, error) {
	if robotInitFunction, ok := robots.Robots[command]; ok {
		return robotInitFunction(), nil
	}

	return nil, fmt.Errorf("unknown robot: %s", command)
}
