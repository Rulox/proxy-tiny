FROM golang:1.10.0 AS builder

WORKDIR /go/src/proxy
COPY . /go/src/proxy

# Install go dependencies
RUN rm -rf /go/src/proxy/vendor && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -v -update && \
    dep ensure -v

# Compiling the binary without dynamic links
RUN go build -tags static main.go

# Final image
FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=builder /go/src/proxy/main /main

CMD ["./main"]
