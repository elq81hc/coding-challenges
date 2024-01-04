# Build target
build:
	go build -o bin/ccwc main.go

# Clean target
clean:
	rm -r bin/

test:
	@go test -v ./...

benchmark:
	@go test -bench=. -benchmem ./...