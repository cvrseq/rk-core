
build:
	@echo "Building Go binary..."
	mkdir -p ./build
	CGO_ENABLED=0 GOOS=linux go build -o ./build/app ./cmd/main.go

clean:
	@echo "Cleaning Go binary files..."
	rm -rf ./build
