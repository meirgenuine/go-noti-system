# Go Notification System

This repository contains a Go-based notification system. It consists of a WebSocket server for sending notifications, a mobile client for receiving and forwarding them, and a gRPC server for the final reception.

## Components

- `websockets-server`: A WebSocket server running on `localhost:8568/ws`.
- `mobile-client`: A mobile client connecting to the WebSocket server and forwarding notifications to a gRPC server.
- `grpc-server`: A gRPC server running on `localhost:50051`, receiving notifications from the mobile client.

## Installation and Setup

1. Clone this repository.
2. Install Go and gomobile.
3. Run the Makefile to build all components.

`make build-all`

## Note

Make sure to install all dependencies to allow proper build for iOS and Android

## License

This project is under the MIT License. See [LICENSE](LICENSE) for details.
