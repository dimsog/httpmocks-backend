FROM golang:1.25.5 AS go-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o httpmocks ./cmd/httpmocks
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.23

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=go-builder /app/httpmocks .
COPY --from=go-builder /go/bin/goose /usr/local/bin/goose
COPY --from=go-builder /app/database/migrations ./database/migrations

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8000

CMD ["./httpmocks"]
