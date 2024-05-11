FROM golang:1.21

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .

RUN go build -o /app/webServerApp

EXPOSE 8000

CMD ["/app/webServerApp"]