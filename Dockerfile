# Use the official Golang image from the Docker Hub
FROM golang:latest

# Create a directory for the application inside the Docker image
WORKDIR /app

# Copy the Go files from your project into the current directory inside the Docker image
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go bot
RUN go build -o main .

# Command to run the bot
CMD ["./main"]
