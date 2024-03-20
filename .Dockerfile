FROM golang:1.22.1-alpine3.19

WORKDIR /app
COPY ./healthcheck ./healthcheck

CMD [ "./healthcheck" ]