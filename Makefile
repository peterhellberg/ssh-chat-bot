BINARY = ssh-chat-bot

all: $(BINARY)

**/*.go:
	go build ./...

$(BINARY): **/*.go *.go
	go build -ldflags "-X main.buildCommit=`git rev-parse --short HEAD`" .

build: $(BINARY)

clean:
	rm -f $(BINARY)
