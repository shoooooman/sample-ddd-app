FROM golang:latest

WORKDIR /go/src/work

CMD ["go", "run", "main.go"]
