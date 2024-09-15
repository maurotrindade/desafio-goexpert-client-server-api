FROM golang:latest AS server_builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o server .


# FROM scratch - possivelmente não funciona por não ter CGO
FROM debian
WORKDIR /app
COPY --from=server_builder /app/server .
COPY --from=server_builder /app/.env .
CMD ["./server"]
# CMD ["tail", "-f", "/dev/null"]
