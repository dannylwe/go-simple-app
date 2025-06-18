# Dockerfile

# --- Builder Stage ---
# Use a specific Go version for reproducibility.
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum to download dependencies.
COPY go.mod ./

# Download dependencies. This is done in a separate layer to leverage Docker's caching.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go application.
# -o /app/main specifies the output file.
# -ldflags "-w -s" strips debugging information, reducing binary size.
# CGO_ENABLED=0 disables Cgo, creating a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# --- Final Stage ---
# Use a minimal base image for the final container to reduce the attack surface.
FROM alpine:latest

# Set the working directory.
WORKDIR /app

# Copy the built binary from the builder stage.
COPY --from=builder /app/main .

# Expose the port the application runs on.
EXPOSE 9002

# Set the user to a non-root user for security.
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Define the command to run the application.
ENTRYPOINT ["/app/main"]
