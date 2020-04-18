package herder

import (
	"fmt"
	"strings"

	"github.com/peterhellberg/ssh-chat-bot/robots"
)

var (
	errWrongUser             = fmt.Errorf("wrong user")
	errNotEnoughFieldsInLine = fmt.Errorf("not enough fields in line")
)

func parsePublicCommand(line, user string) (*robots.Command, error) {
	fields := strings.Fields(line)

	if len(fields) < 3 {
		return nil, errNotEnoughFieldsInLine
	}

	if strings.TrimRight(fields[1], ":") != user {
		return nil, errWrongUser
	}

	args := []string{}

	if len(fields) > 3 {
		for _, f := range fields[3:] {
			args = append(args, strings.TrimRight(f, "\a"))
		}
	}

	return &robots.Command{
		From:    strings.TrimRight(fields[0], ":"),
		Command: strings.TrimRight(fields[2], "\a"),
		Args:    args,
	}, nil
}

func parsePrivateCommand(line string) (*robots.Command, error) {
	fields := strings.Fields(strings.TrimPrefix(line, "[PM from "))

	if len(fields) < 2 {
		return nil, errNotEnoughFieldsInLine
	}

	args := []string{}

	if len(fields) > 2 {
		for _, f := range fields[2:] {
			args = append(args, strings.TrimRight(f, "\a"))
		}
	}

	return &robots.Command{
		From:    strings.TrimRight(fields[0], "]"),
		Command: strings.TrimRight(fields[1], "\a"),
		Args:    args,
	}, nil
}
