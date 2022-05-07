lint:
	go vet ./...
	staticcheck ./... 

test:
	ginkgo -r

install: lint
	go install .
