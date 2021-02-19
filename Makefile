GO = go
PREFIX ?= /usr/local/bin

all: server client

.PHONY: install
install:
	mv multiplexor-server $(PREFIX)
	mv multiplexor-client $(PREFIX)

.PHONY: server
server:
	$(GO) build -o multiplexor-server ./cmd/multiplexor-server

.PHONY: server-windows
server-windows:
	GOOS=windows $(GO) build -o multiplexor-server.exe ./cmd/multiplexor-server

.PHONY: client
client:
	$(GO) build -o multiplexor-client ./cmd/multiplexor-client

.PHONY: client-windows
client-windows:
	GOOS=windows $(GO) build -o multiplexor-client.exe ./cmd/multiplexor-client

.PHONY: clean
clean:
	rm -f multiplexor-server multiplexor-server.exe
	rm -f multiplexor-client multiplexor-client.exe
