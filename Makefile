GOBUILD=go build
GOTEST=go test 


.PHONY: all build server client docker-build

all: test build

build: server client

server: 
	$(GOBUILD) -o ./bin/server ./server

client: 
	$(GOBUILD) -o ./bin/client ./client

test:
	$(GOTEST) ./...

docker-build:
	docker build -t time-server ./server/

