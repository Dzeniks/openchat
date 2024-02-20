# Use the official Go image as the base image
FROM golang:1.21.5-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# RUN go build -o server .

# Expose the port on which the server will listen
EXPOSE 8080

# Set the command to run the server when the container starts
CMD ["go", "run", "main.go"]
