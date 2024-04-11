.PHONY: test
test:
	go test ./...

.PHONY: mod/download
mod/download:
	go mod download

.PHONY: mod/tidy
mod/tidy:
	go mod tidy
