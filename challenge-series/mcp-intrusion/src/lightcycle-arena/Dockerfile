FROM golang:alpine AS builder

COPY ./main.go .

RUN go build -o svc main.go

FROM alpine:latest

COPY --from=builder /go/svc /svc
COPY ./README.md /README.md

CMD ["./svc"]
