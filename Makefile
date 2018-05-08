SHELL=bash

bin/client: client/app.go
	go build -o bin/client ./client

bin/server: server/app.go
	go build -o bin/server ./server

bin/linux/server: server/app.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/server ./server

push/linux/server:
	cd ./infra && scp ../bin/linux/server "ubuntu@$$(terraform output ip_address):/home/ubuntu/server"

ssh/linux/server:
	cd ./infra && ssh "ubuntu@$$(terraform output ip_address)"

infra/teardown:
	cd ./infra && terraform destroy

build: bin/client bin/server bin/linux/server

.PHONY: build
