# Use Ubuntu 22.04 as base image
FROM ubuntu:22.04

# Set environment variables to avoid interactive prompts
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=UTC

# Install system dependencies
RUN apt-get update && apt-get install -y \
    git \
    cmake \
    build-essential \
    pkg-config \
    golang-go \
    libasound2-dev \
    && rm -rf /var/lib/apt/lists/*

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=$PATH:/go/bin
ENV CGO_ENABLED=1

# Create working directory
WORKDIR /app

# Copy the project files
COPY . .

# Build ggwave library
RUN make build

# Install ggwave library
RUN make install

# Run Go tests to verify the build
RUN go test .

# Download dependencies and build the example
RUN go mod download
RUN go mod tidy
RUN cd examples/encode-decode && go build -o encode-decode main.go

# Set the default command
CMD ["bash"] 