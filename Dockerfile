FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /build

COPY src/* .

RUN go mod download && go mod verify
RUN go build -v -o /build/healthcheck

FROM alpine:3.19

RUN apk add --no-cache tzdata

COPY --from=builder /build/healthcheck /usr/local/bin/healthcheck

VOLUME /config
ENV CONFIG_FILE=/config/config.yml

CMD ["healthcheck"]