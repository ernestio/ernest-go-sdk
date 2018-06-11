install:
	go install -v

test:
	go test --cover -v $$(go list ./... | grep -v /vendor/)

cover:
	go test -coverprofile cover.out

deps:
	go get -u gopkg.in/yaml.v2
	go get -u github.com/r3labs/sse
	go get -u github.com/r3labs/diff
	go get github.com/mitchellh/mapstructure
	go get github.com/gorilla/websocket

dev-deps: deps
	go get -u github.com/stretchr/testify
	go get -u github.com/golang/lint/golint
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/ernestio/crypto
	go get -u github.com/r3labs/broadcast

lint:
	golint ./...
	go vet ./...
