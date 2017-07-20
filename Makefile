test:
	go test -v -race $(shell go list ./... | grep -v /vendor/)
