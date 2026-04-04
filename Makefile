
build:
	@echo "Building Go binary..."
	mkdir -p ./build
	go build -o ./build/app ./cmd/main.go

clean:
	@echo "Cleaning Go binary files..."
	rm -rf ./build
