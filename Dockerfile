FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /build

COPY src/* .

RUN go mod download && go mod verify
RUN go build -v -o ./healthcheck

FROM alpine:3.19

COPY - FROM=builder /build/healthcheck /usr/local/bin/
COPY config.yml /config/config.yml

ENV CONFIG_FILE=/config/config.yml
CMD ["app"]