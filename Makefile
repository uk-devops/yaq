lint:
	go vet ./...
	staticcheck ./...

test:
	ginkgo -r --label-filter='!azure'

azure-test:
	ginkgo -r --label-filter='azure' --slow-spec-threshold 30s

install:
	go install .
