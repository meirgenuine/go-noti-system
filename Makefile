# Makefile
# Specific features to manage build processes

# Build Android binaries for mobile-client
build-android-mobile-client:
	@echo "Building mobile-client APK for Android..."
	@cd mobile-client && gomobile build -target=android -o ../build/mobile-client.apk
	@echo "Build completed."

# Build iOS binaries for mobile-client
build-ios-mobile-client:
	@echo "Building mobile-client IPA for iOS..."
	@cd mobile-client && gomobile build -target=ios -o ../build/mobile-client.ipa
	@echo "Build completed."

# Build Android binaries for grpc-server
build-android-grpc-server:
	@echo "Building grpc-server APK for Android..."
	@cd grpc-server && gomobile build -target=android -o ../build/grpc-server.apk
	@echo "Build completed."

# Build iOS binaries for grpc-server
build-ios-grpc-server:
	@echo "Building grpc-server IPA for iOS..."
	@cd grpc-server && gomobile build -target=ios -o ../build/grpc-server.ipa
	@echo "Build completed."

# Build everything
build-all: build-android-mobile-client build-ios-mobile-client build-android-grpc-server build-ios-grpc-server

