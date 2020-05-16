BINARY=build/app

build:
	go build ./src -o $(BINARY)

test:
	go test

clean:
	go clean
	rm -f $(BINARY)

run:
	go build -o $(BINARY)
	./build/app
