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

.PHONY: unit-tests
clean:
	@rm -rf ./bin
	@# mkdir -p ./bin
	@echo $@

.PHONY: compose-run
compose-run: compose-build compose-up
	@docker-compose up
	@echo $@
	
.PHONY: compose-up
compose-up: 
	@docker-compose up
	@echo $@
	
.PHONY: compose-build
compose-build: 
	@docker-compose build
	@echo $@
