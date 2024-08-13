FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from GoReleaser
COPY labeled-storage .

# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the executable
CMD ["./labeled-storage"]
