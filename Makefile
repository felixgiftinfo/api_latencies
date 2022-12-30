.PHONY: run
run:
	@go run main.go
	@echo $@

.PHONY: run-build
run-build: build ; @./bin/api_latency_build
	@echo $@

.PHONY: build
build: clean 
	@go build -o bin/api_latency_build
	@echo $@

.PHONY: air
air:
	@air

.PHONY: unit-tests
unit-tests:
	@go test ./...

clean:
	@rm -rf ./bin
	@# mkdir -p ./bin
	@echo $@
all: 
	@echo "This make line will not be printed"
	echo "But this will"