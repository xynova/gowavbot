#!/bin/bash

# Build the Docker image
echo "Building gogwave Docker image..."
docker build -t gogwave:latest .

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "✅ Docker image built successfully!"
    echo ""
    echo "To run the container:"
    echo "  docker run -it gogwave:latest"
    echo ""
    echo "To run the example:"
    echo "  docker run -it gogwave:latest /app/examples/encode-decode/encode-decode"
    echo ""
    echo "To run tests:"
    echo "  docker run -it gogwave:latest go test ."
else
    echo "❌ Docker build failed!"
    exit 1
fi 