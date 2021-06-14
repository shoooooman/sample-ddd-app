FROM golang:latest

WORKDIR /go/src/work

COPY go.mod .

RUN go mod download

CMD ["go", "run", "main.go"]
