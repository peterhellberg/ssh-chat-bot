package robots

import "testing"

func TestAboutBot(t *testing.T) {
	bot, err := GetRobot("about")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd := &Command{}

	if got, want := bot.Run(cmd), aboutRawurl; got != want {
		t.Fatalf("bot.Run(cmd) = %q, want %q", got, want)
	}

	if got, want := bot.Description(), ""; got != want {
		t.Fatalf("bot.Description() = %q, want %q", got, want)
	}
}
