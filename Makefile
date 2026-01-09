build:
	go build -o bin/helix

run: build
	./bin/helix
