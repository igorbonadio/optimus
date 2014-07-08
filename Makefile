.PHONY: test
test:
	@echo "Testing lbfgs"
	@ginkgo test/lbfgs

.PHONY: fmt
fmt:
	gofmt -w $(shell find . -name '*.go')