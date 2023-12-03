# Start from the latest golang base image
FROM golang:alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gps-data-service .

# Start a new stage from scratch
FROM scratch  

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/gps-data-service .

# Command to run the executable
CMD ["./gps-data-service"]
