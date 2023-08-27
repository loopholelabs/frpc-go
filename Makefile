.PHONY: all
all: tests

vet:
	go vet ./...

tests: vet
	@go clean -testcache
	go test -p 1 -race  ./... -coverpkg=./spire/pkg/...  -coverprofile cover.out && go tool cover -func=cover.out