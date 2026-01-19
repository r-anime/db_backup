FROM golang:1.25.6-alpine AS builder

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

COPY --from=builder /app/db_backup .
RUN chmod +x ./db_backup

ENTRYPOINT ["./db_backup"]
