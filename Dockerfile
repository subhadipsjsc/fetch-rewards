# Use the official Go image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy all project files to the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application into a binary
RUN go build -o app .

# Set 777 permissions for the app binary
RUN chmod 777 app

# Expose the application port
EXPOSE 8080

# Command to run the compiled binary
CMD ["./app"]
