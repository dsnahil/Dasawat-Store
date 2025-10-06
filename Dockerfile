# ---- build stage ----
FROM golang:1.22-alpine AS build
WORKDIR /app

# Copy only the Go source; we don't depend on a host go.mod
COPY main.go .

# Create a minimal module inside the container so go build succeeds reliably
RUN go mod init productapi && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /product-api ./main.go

# ---- final stage ----
FROM alpine:latest
WORKDIR /
COPY --from=build /product-api /product-api
EXPOSE 8081
ENTRYPOINT ["/product-api"]
