FROM golang:1.22 AS builder
WORKDIR /build
COPY src/* .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o pagemaster

FROM golang:alpine
WORKDIR /app
COPY --from=builder /build/pagemaster /app/pagemaster
COPY utils/pagefile.dat /app/pagefile.dat
COPY memory/memory.go /app/memory.go
EXPOSE 8080
CMD ["/app/pagemaster"]