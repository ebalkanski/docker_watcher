FROM golang:1.22.1

RUN go install github.com/ysmood/kit/cmd/guard@v0.25.11

ADD . /go/src/docker_watcher

WORKDIR /go/src/docker_watcher

ENV TZ=Europe/Sofia

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "/go/bin/guard -w '**/*.go' -- go run ./cmd/..."]
