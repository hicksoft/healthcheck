FROM golang:1.22.1-alpine3.19

WORKDIR /usr/src/app

COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

COPY src/* .
RUN go build -v -o /usr/local/bin/app

CMD ["app"]