include .env

.PHONY: dev push crt

dev:
# go install github.com/cosmtrek/air@latest
# export PATH=$PATH:$(go env GOPATH)/bin
	air

push:
	docker build -t app .
	docker login $(ACR_HOST) -u $(ACR_USER) -p $(ACR_PASSWORD)
	docker tag app $(ACR_HOST)/app:latest
	docker push $(ACR_HOST)/app:latest

crt:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
