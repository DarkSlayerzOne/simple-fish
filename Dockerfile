FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . /app/

ENV PORT 8090

RUN go build

CMD ["./simple-fish-api"]