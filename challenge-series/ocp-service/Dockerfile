FROM golang:1.22 AS build

COPY src /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./...

FROM alpine:latest
COPY --from=build /build/ocp-service /app/ocp-service
EXPOSE 8080
CMD ["/app/ocp-service"]