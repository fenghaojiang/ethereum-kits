.PHONY: build


lint: 
	golangci-lint run ./...



generate: 
	go generate ./contracts/...
	go mod tidy