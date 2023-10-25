.PHONY: build




generate: 
	go generate ./constant
	go mod tidy