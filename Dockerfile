FROM golang:1.22

# Set working directory inside container
WORKDIR /go/src/app

# Copy files
COPY ./cmd /go/src/app/cmd
COPY ./config /go/src/app/config
COPY ./internal /go/src/app/internal
COPY ./static /go/src/app/static
COPY ./utils /go/src/app/utils
COPY go.mod go.sum /go/src/app/

# Move into server directory
WORKDIR /go/src/app/cmd/server

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

# Expose port
EXPOSE 8080

# Run the server
CMD ["./server"]

