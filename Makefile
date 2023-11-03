.PHONY: help lint test doc
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

pre-push: lint test ## Run golang lint and test

lint: ## Run golang lint using docker
	go mod download
	docker run --rm \
		-v ${GOPATH}/pkg/mod:/go/pkg/mod \
 		-v ${PWD}:/app \
 		-w /app \
	    golangci/golangci-lint:v1.55.2 \
	    golangci-lint run -v --modules-download-mode=readonly

test: ## Run tests
	go test ./...

doc: ## Run doc server using docker
	@echo "Doc server runs on http://127.0.0.1:6060"
	docker run --rm \
        -p 127.0.0.1:6060:6060 \
        -v ${PWD}:/go/src/github.com/evsamsonov/trading-timeseries \
        -w /go/src/github.com/evsamsonov/trading-timeseries  \
        golang:latest \
        bash -c "go install golang.org/x/tools/cmd/godoc@latest && /go/bin/godoc -http=:6060"
