FROM golang:1.24.3-alpine AS base

# Install essential dependencies
RUN apk add --no-cache git nodejs npm

# Install Wails CLI globally
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

# BACKEND INSIDE /go-back

# Set the working directory for the backend
WORKDIR /go-back

# Copy only files that define Go dependencies
COPY backend/go.mod backend/go.sum ./
# Download Go dependencies
RUN go mod download

# Copy the rest of the backend source code
COPY backend/ ./

# FRONTEND 
# Copy frontend/master from host to /go-back/frontend
COPY frontend/master ./frontend

# Expose the ports Wails uses in dev mode: 3000 for Go backend and 8080 for Vite frontend
EXPOSE 3000 8080

# When container builds, run Wails in dev mode
CMD ["wails", "dev"]