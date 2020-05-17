BINARY=build/app

build:
	go build ./src -o $(BINARY)

test:
	go test ./tests/...

clean:
	go clean
	rm -f $(BINARY)

run:
	go build -o $(BINARY)
	./$(BINARY)
