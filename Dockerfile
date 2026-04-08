FROM golang:1.26.2-alpine AS builder

RUN apk add --no-cache git

ENV APP_HOME=/app
WORKDIR $APP_HOME

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .

RUN go build -o db_backup .

FROM alpine:latest AS runtime

ENV APP_HOME=/app
WORKDIR $APP_HOME

RUN apk add --no-cache docker-cli && rm -rf /var/cache/apk/*

COPY --from=builder /app/db_backup .
RUN chmod +x ./db_backup

ENTRYPOINT ["./db_backup"]
