MAIN_FILE_NAME=main
BINARY_NAME=server

dev:
	@echo "-->Run dev mode"
	go run ./cmd/$(MAIN_FILE_NAME).go

test:
	echo "-->Run test"
	go test -v ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

clean:
	@echo "-->Clean"
	go clean
	rm -rf test.db
	rm -rf bin

dep: clean
	@echo "-->Download dependencies"
	go mod download
	go mod verify
	go mod tidy

build: dep
	@echo "==>Building binary"
	go build -o bin/$(BINARY_NAME) -v ./cmd/$(MAIN_FILE_NAME).go

run: build
	@echo "==>Run binary"
	./bin/$(BINARY_NAME)