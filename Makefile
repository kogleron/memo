
build:
	go build -o ./bin/memo ./cmd/memo

run: build
	./bin/memo
