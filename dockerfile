# Use an official Golang runtime as a parent image
FROM golang:latest AS builder

# Set the working directory to /app
WORKDIR /cmd

# Copy the current directory contents into the container at /app
COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main

COPY --from=builder /cmd/main ./

# Download and install any required dependencies
RUN go mod download

FROM alpine:latest

# Expose port 8000 for incoming traffic
EXPOSE 8000

# Define the command to run the app when the container starts
CMD ["./cmd/main.go", "--port", "8000"]