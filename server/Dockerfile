FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o server .


FROM scratch
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/.env .
CMD ["./server"]
