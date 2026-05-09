
build:
	@echo "Building Go binary..."
	mkdir -p ./bin
	go build -o ./bin/app ./cmd/main.go

clean:
	@echo "Cleaning Go binary files..."
	rm -rf ./bin
