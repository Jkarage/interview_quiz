build: 
	@go build -o ./bin/asessment

run: build
	@./bin/asessment

test: 
	@go test -v ./...