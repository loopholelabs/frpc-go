.PHONY: proto
proto: install
	protoc --go-frpc_out=./pkg/generator ./examples/test/test.proto

.PHONY: install
install:
	go install ./protoc-gen-go-frpc

.PHONY: test
test: proto
	go test -v ./...
