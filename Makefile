.PHONY: dev

dev:
# go install github.com/cosmtrek/air@latest
# export PATH=$PATH:$(go env GOPATH)/bin
	air

crt:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
