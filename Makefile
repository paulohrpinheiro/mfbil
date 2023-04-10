COVERAGE_FILE=coverage.out
test:
	go test -v ./... -coverprofile ${COVERAGE_FILE} -covermode count

coverage: test
	go tool cover -func ${COVERAGE_FILE}

.PHONY: test coverage
