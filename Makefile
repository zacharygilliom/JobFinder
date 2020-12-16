build:
	go build -o bin/findjob cmd/main.go

run:
	bin/findjob

all: build run