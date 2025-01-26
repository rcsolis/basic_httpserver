BINARY_NAME_SERVER=server

dev:
	@echo "-->Run dev mode"
	go run ./cmd/$(BINARY_NAME_SERVER).go

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
	go build -o bin/ -v ./cmd/server/$(BINARY_NAME_SERVER).go

run: build
	@echo "==>Run binary"
	./bin/$(BINARY_NAME_SERVER)