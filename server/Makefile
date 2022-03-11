LD_FLAGS = -ldflags "-X main.mode=release"
BUILDDIR = build
NAME = chat-server

linux-amd64: *.go
	GOOS=linux GOARCH=amd64 go build $(LD_FLAGS) -o $(BUILDDIR)/$(NAME)-linux-amd64

linux-arm64: *.go
	GOOS=linux GOARCH=arm64 go build $(LD_FLAGS) -o $(BUILDDIR)/$(NAME)-linux-arm64

linux-arm: *.go
	GOOS=linux GOARCH=arm go build $(LD_FLAGS) -o $(BUILDDIR)/$(NAME)-linux-arm

darwin-amd64: *.go
	GOOS=darwin GOARCH=amd64 go build $(LD_FLAGS) -o $(BUILDDIR)/$(NAME)-darwin-amd64

darwin-arm64: *.go
	GOOS=darwin GOARCH=arm64 go build $(LD_FLAGS) -o $(BUILDDIR)/$(NAME)-darwin-arm64

windows-amd64: *.go
	GOOS=windows GOARCH=amd64 go build $(LD_FLAGS) -o $(BUILDDIR)/$(NAME)-windows-amd64.exe

all: linux-amd64 linux-arm64 linux-arm darwin-amd64 darwin-arm64 windows-amd64
