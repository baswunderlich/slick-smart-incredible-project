# Use an official Node.js image as the base image
FROM node:18-slim

# Install Go
RUN apt-get update && apt-get install -y wget \
    && wget https://go.dev/dl/go1.21.8.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.8.linux-amd64.tar.gz \
    && rm go1.21.8.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV PATH=$PATH:/go/bin
ENV GOPATH=/go

# Set the working directory for the application
WORKDIR /app

# Copy the package.json and package-lock.json (if available)
COPY . .

# Install wails
RUN rm -r ./frontend/node_modules
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

#Install gcc
RUN apt-get update && \
    && apt-get install -y \
    && build-essential \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
    && apt-get install -y libgtk-3-dev

#Install pkg-config
RUN apt update \
    && apt -y install pkg-config

# Command to run your Go application or Node.js application
# You can modify this to run your specific Go or Node.js application
CMD ../go/bin/wails dev

# Expose a port (if needed)
EXPOSE 80
