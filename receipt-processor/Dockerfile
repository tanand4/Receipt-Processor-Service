# Use the official Golang image to create a build artifact.
FROM golang:1.16 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o receipt-processor ./cmd/server

# Use the official Alpine image for a lean production container.
FROM alpine:3.13
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/receipt-processor /receipt-processor

# Run the binary.
ENTRYPOINT ["/receipt-processor"]
