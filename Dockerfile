# Stage 1: Build the application
FROM golang:1.23 AS builder

# Enable CGO and set the target OS
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build a static binary
RUN go build -a -installsuffix cgo -o main .

# Stage 2: Create a lightweight image to run the application
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose application port (change this if your app uses a different port)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
