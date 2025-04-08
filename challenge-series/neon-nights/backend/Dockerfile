FROM golang:1.22 AS builder
WORKDIR /build
COPY src/* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o miami-first

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/miami-first /app/miami-first
COPY people.json /app/people.json
CMD ["/app/miami-first"]