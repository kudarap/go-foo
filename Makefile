APPNAME=foosvc
LDFLAGS="-X main.vTag=`cat VERSION` \
		-X main.vCommit=`git rev-parse HEAD` \
		-X main.vBuilt=`date -u +%s`"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

all: test build

test:
	go test -v -cover -race ./...

run: build
	./$(APPNAME)

run-server: build
	./$(APPNAME) server

run-worker: build
	./$(APPNAME) worker

build:
	CGO_ENABLED=0 go build -v -ldflags=$(LDFLAGS) ./cmd/$(APPNAME)

local-dbs:
	docker run --rm --name foo-postgres -d -v "$(PWD)/.localdata/postgres":/var/lib/postgresql/data -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -p 5432:5432 postgres:16.2