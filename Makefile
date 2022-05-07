lint:
	go vet ./...
	staticcheck ./...

test:
	ginkgo -r --label-filter='!azure'

azure-test:
	ginkgo -r --label-filter='azure'

install: lint
	go install .
