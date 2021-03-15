run:
		go run cmd/api/main.go

fmt:
		go fmt ./...
	 	strictgoimports -w -local "myapp" .
