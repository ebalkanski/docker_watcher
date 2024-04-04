# Overview

Simple HTTP server running into a Docker container with a watcher that restarts it when code is changed.

## Prerequisites

* Install `go`
* Set `GOPATH`
* Install `docker`
* Clone project into $GOPATH/src/docker_watcher

## Run with Docker

Build the Docker image.
```bash
docker build -t docker_watcher -f Dockerfile .
```

Start the `docker_watcher` container.
```bash
docker run -d --rm -p 8080:8080 -v $GOPATH/src/docker_watcher:/go/src/docker_watcher --name docker_watcher docker_watcher
```

The command will start one Docker container with a service listening
for requests on port 8080 and accessible via browser on `http://localhost:8080`.

There is a code watcher running in the container, so if you make a code change the service will be restarted
automatically and you will be able to see the changes immediately in the browser.

You can see the logs of the running container by executing:
```bash
docker logs -f docker_watcher
```

Test it by opening [http://localhost:8080](http://localhost:8080) in a browser.
