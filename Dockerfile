# Stage 1: Build the Go application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy the Go application source code into the container
COPY *.go ./
COPY go.mod .
COPY go.sum .
COPY needed/* needed/

# Build the Go application
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o myCV

# Stage 2: Create a minimal Docker image
FROM alpine:3.14

WORKDIR /app
COPY components/* components/
COPY icons/* icons/
COPY pages/* pages/
COPY data/* data/
# Copy the binary from the builder stage into the final image
COPY --from=builder /app/myCV .

# Expose port 8080
EXPOSE 8080

# Define the command to run the application
CMD ["./myCV"]
