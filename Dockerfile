# Stage 1: The 'build' stage
# We use a specific Go version to build our application.
FROM golang:1.21-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies first
# This is a Docker layer caching optimization
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application.
# -o /product-api specifies the output file name.
# CGO_ENABLED=0 creates a static binary that doesn't depend on system libraries.
RUN CGO_ENABLED=0 go build -o /product-api .

# Stage 2: The 'final' stage
# We use a minimal 'scratch' image which is completely empty.
FROM scratch

# Set the working directory
WORKDIR /

# Copy the built binary from the 'build' stage
COPY --from=build /product-api /product-api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/product-api"]