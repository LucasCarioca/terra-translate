.PHONY := all

test:
	go test ./pkg/terraform -v -covermode=count -coverprofile=coverage.out

lint:
	golint ./...

vet:
	go vet ./...