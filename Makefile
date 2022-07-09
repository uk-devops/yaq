lint:
	go vet ./...
	staticcheck ./...

test:
	ginkgo -r --label-filter='!azure'

azure-test:
	ginkgo -r --label-filter='azure' --slow-spec-threshold 30s

install:
	go install .

darwin:
	$(eval GOOS=darwin)

linux:
	$(eval GOOS=linux)

windows:
	$(eval GOOS=windows)
	$(eval EXE_SUFFIX=.exe)

arm64:
	$(eval GOARCH=arm64)

amd64:
	$(eval GOARCH=amd64)

build:
	$(if ${VERSION}, , $(eval VERSION=latest))
	mkdir -p tmp
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o tmp/yaq${EXE_SUFFIX} .
	zip -j tmp/yaq_${GOOS}_${GOARCH}_${VERSION}.zip tmp/yaq${EXE_SUFFIX}
	rm tmp/yaq${EXE_SUFFIX}
