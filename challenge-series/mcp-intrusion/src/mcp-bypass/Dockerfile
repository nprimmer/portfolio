# Step 1: Build the Go binary in a builder container
FROM golang:1.22 AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY main.go ./

# Build the Go binary
RUN go build -o app .

# Step 2: Build the final container from Alpine
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the binary from the builder container


# Copy everything from the current directory to the new directory
COPY .env .
RUN rm -f .env
COPY *.db .
RUN rm -f *.db

# Expose the port the application will run on
EXPOSE 8080
COPY --from=builder /app/app .
# Command to run the binary
CMD ["./app"]
