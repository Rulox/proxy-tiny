FROM golang:1.9.0

ARG ssh_prv_key
ARG ssh_pub_key

WORKDIR /go/src/http_server
COPY . /go/src/http_server

RUN mkdir /root/.ssh && \
    go get -u github.com/golang/dep/cmd/dep && \
    echo '[url "git@gitlab.s73cloud.com"]\n    insteadOf = "https://gitlab.s73cloud.com"' >> /root/.gitconfig && \
    echo 'Host gitlab.s73cloud.com\n    Port 13579\n    User git\nStrictHostKeyChecking no' >> /root/.ssh/config && \
    dep ensure -v -update && \
    dep ensure -v

RUN go build -tags static main.go

CMD ["/go/src/http_server/main"]