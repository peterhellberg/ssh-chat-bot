package herder

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/peterhellberg/ssh-chat-bot/robots"
	"golang.org/x/crypto/ssh"
)

type Herder struct {
	user    string
	owner   string
	addr    string
	delay   time.Duration
	check   time.Duration
	verbose bool
	active  bool
}

func New(options ...Option) *Herder {
	h := &Herder{}

	for _, option := range options {
		option(h)
	}

	return h
}

func (h *Herder) validate() error {
	switch {
	case h.user == "":
		return fmt.Errorf("missing user")
	case h.owner == "":
		return fmt.Errorf("missing owner")
	case h.addr == "":
		return fmt.Errorf("missing addr")
	default:
		return nil
	}
}

// Run the herder controlling the robots
func (h *Herder) Run() error {
	if err := h.validate(); err != nil {
		return err
	}

	session, err := newSession(h.addr, h.user)
	if err != nil {
		return err
	}
	defer session.Close()

	if err := session.Setenv("TERM", "bot"); err != nil {
		return err
	}

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

	if err := session.RequestPty("xterm", 80, 40, ssh.TerminalModes{}); err != nil {
		return err
	}

	go func() {
		time.Sleep(h.delay)
		if h.owner != "" {
			in.Write([]byte("/msg " + h.owner + " Now active\r\n"))
		}
		h.active = true
	}()

	go func() {
		for {
			in.Write([]byte("/motd\r\n"))
			time.Sleep(h.check)
		}
	}()

	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return err
		}

		if !strings.Contains(line, ": "+h.user+": ") {
			continue
		}

		cmd, err := parseCommand(line, h.user)
		if err == nil {
			if h.active {
				fmt.Printf("%#v\n", cmd)
			}

			robot, err := robots.GetRobot(cmd.Command)
			if err != nil {
				continue
			}

			if !h.active {
				continue
			}

			if response := robot.Run(cmd); response != "" {
				reply(in, fmt.Sprintf("%s %s", cmd.From, response))
			}
		}
	}

	return fmt.Errorf("ERROR")
}

func reply(in io.WriteCloser, s string) {
	in.Write([]byte(s + "\r\n"))
}

func newSession(addr, user string) (*ssh.Session, error) {
	conn, err := dial(addr, user)
	if err != nil {
		return nil, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func dial(addr, user string) (*ssh.Client, error) {
	key, err := makeKey()
	if err != nil {
		return nil, err
	}

	return ssh.Dial("tcp", addr, clientConfig(user, key))
}

func clientConfig(user string, key ssh.Signer) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
}

// makeKey makes a SSH signer key
func makeKey() (ssh.Signer, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		return nil, err
	}

	return ssh.NewSignerFromKey(key)
}
