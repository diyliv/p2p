# Secure P2P Chat with RSA Encryption
A lightweight peer-to-peer chat application written in Go, featuring end-to-end RSA encryption and support for multiple concurrent connections. Each peer acts as a both a server (listening for incoming connections) and a client (able to connect to other peers). 

## Features 
- **Multi-connection support** - each peer can maintain simultaneous connections with multiple remote peers
- **End-to-end encryption** - every message is encrypted with RSA-2048 and the remote peer's public key