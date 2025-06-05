# Use the official Golang image as base for building and running the app
FROM golang:1.24.3-alpine

# Install essential dependencies
RUN apk add --no-cache git nodejs npm

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

# Expose the ports used during development: 3000 for the Go API and 8080 for the Vite dev server
EXPOSE 3000 8080

# Default command launches the development server
CMD ["wails", "dev"]
