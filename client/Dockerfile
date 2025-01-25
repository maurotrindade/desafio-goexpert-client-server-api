FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o client .


FROM debian
WORKDIR /app
COPY --from=builder /app/client .
COPY --from=builder /app/.env .
CMD ["./client"]
# CMD ["tail", "-f", "/dev/null"]
