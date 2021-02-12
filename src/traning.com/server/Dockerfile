FROM golang:latest as builder

WORKDIR /go/src/traning.com/server

COPY . /go/src/traning.com/server

RUN go build .

EXPOSE 8080

CMD ["/go/src/traning.com/server/server"]
