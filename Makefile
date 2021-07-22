fmt: ## run go fmt against code
	go fmt ./internal/... ./cmd/... 

vet: ## run go vet against code
	go vet ./internal/... ./cmd/... 

build: ## build binary called geasy
	go build -o geasy .

clean: ## clean up workspace
	rm -f geasy
	go clean -testcache

install: build ## install binary to gopath bin folder
	mv geasy ${GOPATH}/bin/geasy

.PHONY: clean vet fmt build install clean