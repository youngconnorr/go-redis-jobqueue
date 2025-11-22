# Use official Go image as base
FROM golang:1.25

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build the Go program into an executable
RUN go build -o taskqueue ./cmd/taskqueue

# Run the executable when container starts
CMD ["./taskqueue"]