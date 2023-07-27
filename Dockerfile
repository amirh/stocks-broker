# Use an official Golang runtime as the base image
FROM golang:1.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o stocks-broker

# Use a minimal base image to reduce the container size
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/stocks-broker .

# Expose the gRPC port
EXPOSE 50051

# Command to run the service
CMD ["./stocks-broker"]
