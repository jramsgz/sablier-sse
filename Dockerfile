# Start from the official Go base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .
COPY go.mod .

# Build the Go application
RUN go build -o sse-server .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the executable
CMD ["./sse-server"]
