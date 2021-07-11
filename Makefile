.PHONY: all

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v

build: fmt vet test

publish:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/amalfra/oexec@v${VERSION} || true
