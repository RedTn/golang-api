start:
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
