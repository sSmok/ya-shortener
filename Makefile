LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml
lint-fix:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml --fix

test-all:
	make test-shortenertestbeta
	make test-statictest
test-statictest:
	go vet -vettool=$(LOCAL_BIN)/statictest ./...
test-shortenertestbeta:
	$(LOCAL_BIN)/shortenertestbeta -test.v -test.run=^TestIteration1$$ -binary-path=/Users/admin/go_projects/yandex/ya-shortener/cmd/shortener/shortener