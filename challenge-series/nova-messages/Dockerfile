FROM golang:1.22 AS builder
WORKDIR /build
COPY main.go main.go
COPY messages.go messages.go
COPY go.* .
RUN go mod download
RUN GOOS=linux go build -o nova-messages

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /build/nova-messages /app/nova-messages
CMD ["/app/nova-messages"]