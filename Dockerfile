FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.21

RUN addgroup -S appgroup && \
    adduser -S -G appgroup appuser

WORKDIR /app

COPY --from=builder /app/server /app/server

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s \
  --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/health || exit 1

CMD ["/app/server"]
