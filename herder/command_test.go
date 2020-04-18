package herder

import (
	"reflect"
	"testing"

	"github.com/peterhellberg/ssh-chat-bot/robots"
)

func TestParseCommand(t *testing.T) {
	for _, tt := range []struct {
		line string
		user string
		want *robots.Command
		err  error
	}{
		{
			"",
			"",
			nil,
			errNotEnoughFieldsInLine,
		},
		{
			"peter: wrong-user: eval 5 * 110\r",
			"ssh-chat-bot",
			nil,
			errWrongUser,
		},
		{
			"peter: ssh-chat-bot: eval 5 * 110\r",
			"ssh-chat-bot",
			&robots.Command{
				From:    "peter",
				Command: "eval",
				Args: []string{
					"5",
					"*",
					"110",
				},
			},
			nil,
		},
	} {
		got, err := parseCommand(tt.line, tt.user)
		if err != tt.err {
			t.Fatalf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %#v, want %#v", got, tt.want)
		}
	}
}
