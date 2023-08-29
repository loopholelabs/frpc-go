FROM golang as builder

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

RUN go install github.com/loopholelabs/frpc-go/protoc-gen-go-frpc@v0.7.3

# Note, the Docker images must be built for amd64. If the host machine architecture is not amd64
# you need to cross-compile the binary and move it into /go/bin.
RUN bash -c 'find /go/bin/${GOOS}_${GOARCH}/ -mindepth 1 -maxdepth 1 -exec mv {} /go/bin \;'

FROM scratch

# Runtime dependencies
LABEL "build.buf.plugins.runtime_library_versions.0.name"="github.com/loopholelabs/frpc-go"
LABEL "build.buf.plugins.runtime_library_versions.0.version"="v0.7.3"

COPY --from=builder /go/bin /

ENTRYPOINT ["/protoc-gen-go-frpc"]