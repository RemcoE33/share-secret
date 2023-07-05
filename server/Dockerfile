# Building app
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage so there is a small binary withoud go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080 
CMD ["/app/main"]
