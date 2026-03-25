# Stage 1: Build
FROM golang:1.25.7 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Copy source code (must be before go mod tidy)
COPY . .

# Download dependencies
RUN go mod download

# Tidy modules
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o learninggo .

# Check if binary was created
RUN ls -la /app/learninggo || echo "Build failed - binary not found!"

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS (if needed)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/learninggo .

# Check if binary exists
RUN ls -la /root/learninggo || echo "Binary not found!"

# Make binary executable
RUN chmod +x ./learninggo

# Expose ports if needed (optional)
# EXPOSE 8080

# Run the application
CMD ["./learninggo"]
