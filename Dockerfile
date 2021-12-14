FROM golang:alpine3.14 as builder
RUN mkdir /src
ADD . /src/
WORKDIR /src
RUN go build -ldflags "-s -w -X main.version=$(cat VERSION)" -o sfs
EXPOSE 8100
ENTRYPOINT ["/src/sfs", "-d","/src/files"]
