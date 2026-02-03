# Secure P2P Chat with RSA Encryption

## Documentation

This project now has multilingual documentation. Please visit the documentation directory:

### [Go to Documentation →](docs/index.md)

## Project Overview

A lightweight peer-to-peer chat application written in Go, featuring end-to-end RSA encryption and support for multiple concurrent connections. Each peer acts as both a server (listening for incoming connections) and a client (able to connect to other peers).

## Quick Start

```bash
# Clone the repository
git clone https://github.com/diyliv/p2p.git
cd p2p

# Build and run
go build ./cmd/p2pchat
./p2pchat
```


## Development

```bash
# Run tests
go test ./...

# Format code
go fmt ./...

# Build for specific platform
GOOS=linux GOARCH=amd64 go build ./cmd/p2pchat
```
