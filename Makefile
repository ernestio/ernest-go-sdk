install:
	go install -v

test:
	go test -v ./...

cover:
	go test -coverprofile cover.out

deps:
	go get -u gopkg.in/yaml.v2
	go get -u github.com/r3labs/sse
	go get github.com/mitchellh/mapstructure

dev-deps: deps
	go get -u github.com/golang/lint/golint
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/ernestio/crypto

lint:
	golint ./...
	go vet ./...
