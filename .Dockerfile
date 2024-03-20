FROM golang:1.22.1-alpine3.19

COPY /app/healthcheck /app/healthcheck

CMD [ "/app/healthcheck" ]