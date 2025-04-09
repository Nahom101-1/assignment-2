FROM golang:1.23 AS builder

# Set working directory inside container
WORKDIR /go/src/app

# Copy files
COPY . .

# Move into server directory
WORKDIR /go/src/app/cmd/server

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server


# Expose port
EXPOSE 8080

# Run the server
CMD ["./server"]

LABEL stage=builder