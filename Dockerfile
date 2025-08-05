# Build stage for Svelte
FROM node:18-alpine as build

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Build the Svelte app
RUN npm run build

# Production stage with Go
FROM golang:1.21-alpine AS go-build

WORKDIR /app

# Copy Go modules
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy Go source
COPY main.go .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Final stage
FROM alpine:3.18

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Go binary
COPY --from=go-build /app/server .

# Copy the built Svelte app
COPY --from=build /app/dist ./dist

# Expose port
EXPOSE 8080

# Run the server
CMD ["./server"]