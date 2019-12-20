# This dockerfile is based off the official Golang images
# https://hub.docker.com/_/golang

FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
