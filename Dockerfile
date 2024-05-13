FROM golang:1.21

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN go mod init main
RUN go mod tidy
RUN go mod download

COPY *.go .
RUN  go get github.com/sirupsen/logrus
RUN go build -o /app/webServerApp

EXPOSE 8000

CMD ["/app/webServerApp"]