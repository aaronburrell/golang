FROM golang
MAINTAINER Aaron Burrell <aaronburrell@gmail.com>

# Copy the local package files to the container's workspace.
#ADD . /go/src/github.com/golang/example/outyet

RUN go get github.com/aaronburrell/golang/hello
RUN go install github.com/aaronburrell/golang/hello

# Run the hello command by default when the container starts.
ENTRYPOINT /go/bin/hello

# Document that the service listens on port 8080.
EXPOSE 8080
