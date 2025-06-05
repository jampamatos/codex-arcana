# Use the official Golang image as base for building and running the app
FROM golang:1.24.3-bullseye

# Install essential dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y --no-install-recommends \
        git \
        nodejs \
        build-essential \
        pkg-config \
        libgtk-3-dev \
        libwebkit2gtk-4.0-dev \
        libglib2.0-dev \
        libgdk-pixbuf2.0-dev && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Wails CLI globally
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

# BACKEND INSIDE /go-back
# Set the working directory for the backend where Go code lives
WORKDIR /go-back

# Copy only files that define Go dependencies
COPY backend/go.mod backend/go.sum ./
# Download Go dependencies
RUN go mod download

# Copy the rest of the backend source code
COPY backend/ ./

# FRONTEND 
# Copy React source into the expected Wails frontend directory
COPY frontend/master ./frontend

# Expose the ports used during development:
    # 3000: backend Go
    # 5173: Vite (frontend dev server)
    # 34115: Wails DevWatcher
EXPOSE 3000 5173 34115

# Default command launches the development server
#  Vite will be started with --host so the dev server is reachable outside the container
CMD ["wails", "dev"]
