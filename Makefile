LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

test-all:
	make test-shortenertestbeta
	make lint
test-shortenertestbeta:
	$(LOCAL_BIN)/shortenertestbeta -test.v -test.run=^TestIteration7$$ -binary-path=cmd/shortener/shortener -source-path=.
