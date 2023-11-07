.PHONY: build




generate: 
	go generate ./contracts/...
	go mod tidy