build:
	go build -o bin/helix

run: build
	./bin/helix

newrun: build
	./bin/helix --listenaddr :4000 --leaderaddr :3000
