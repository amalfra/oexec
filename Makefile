.PHONY: all

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v

build: fmt vet test

publish:
	cd ../ && go get github.com/amalfra/oexec@${VERSION} || true
