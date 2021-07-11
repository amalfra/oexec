.PHONY: all

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v

build: fmt vet test

publish:
	curl https://proxy.golang.org/github.com/amalfra/oexec/@v/${VERSION}.info
