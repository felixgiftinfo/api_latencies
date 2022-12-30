run:
	go run main.go

air:
	air

build:
	go build -o build/api_latency_build

unit-test:
	go test ./...