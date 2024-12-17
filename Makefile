LOCAL_BIN:=$(CURDIR)/bin

test-all:
	make test-shortenertestbeta
	make test-statictest
test-statictest:
	go vet -vettool=$(LOCAL_BIN)/statictest ./...
test-shortenertestbeta:
	$(LOCAL_BIN)/shortenertestbeta -test.v -test.run=^TestIteration1$$ -binary-path=/Users/admin/go_projects/yandex/ya-shortener/cmd/shortener/shortener