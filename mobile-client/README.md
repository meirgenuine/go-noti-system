# Mobile Client

This is the mobile client part of the Go Notification System. It connects to the WebSocket server to receive notifications, prints the notification and forwards them to a gRPC server.

## Building and Running

Use the following Makefile targets to build the mobile client:

- For Android: `make build-android-mobile-client`
- For iOS: `make build-ios-mobile-client`
