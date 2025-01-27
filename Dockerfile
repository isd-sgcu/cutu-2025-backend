# Build stage
FROM golang:1.22.4-alpine3.20 as builder

# Install required build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/main.go

# Runtime stage
FROM alpine:3.20

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the built binary from the builder stage
COPY --from=builder /app/server ./

# Copy the .env file to the runner stage
COPY .env ./


# Expose the application port
EXPOSE 4000

# Run the application
CMD ["./server"]