# Use Ubuntu 22.04 as the base image
FROM ubuntu:22.04

# Set environment variables to prevent interactive prompts during package installation
ARG DEBIAN_FRONTEND=noninteractive

# Update package list and install necessary tools and dependencies
RUN apt-get update && \
    apt-get install -y \
    git \
    curl \
    gnupg \
    make \
    build-essential \
    pkg-config \
    protobuf-compiler \
    protobuf-c-compiler \
    python3 python3-pip \
    && rm -rf /var/lib/apt/lists/*

# Install Go 1.21 based on the system architecture
RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then \
    GO_ARCH="go1.21.0.linux-amd64.tar.gz"; \
    elif [ "$ARCH" = "arm64" ]; then \
    GO_ARCH="go1.21.0.linux-arm64.tar.gz"; \
    else \
    echo "Unsupported architecture"; exit 1; \
    fi && \
    curl -fsSL "https://go.dev/dl/$GO_ARCH" -o "$GO_ARCH" && \
    tar -C /usr/local -xzf "$GO_ARCH" && \
    rm "$GO_ARCH" && \
    # Install Go language server
    /usr/local/go/bin/go install golang.org/x/tools/gopls@latest

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.54.2

# Set up Go environment variables
ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

# Install Pprotoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Add protoc-gen-go to the PATH
ENV PATH="${PATH}:/root/go/bin"

# Install npm and Node.js v20
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Install Python protobuf compiler
RUN pip3  install "betterproto[compiler]"