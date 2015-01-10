package robots

// Robot is the interface all robots must follow
type Robot interface {
	Run(*Command) string
	Description() string
}

// Command represents the fields in a ssh-chat-bot command
type Command struct {
	From    string
	Command string
	Args    []string
}
