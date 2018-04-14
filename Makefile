GOOS ?= linux
GOARCH ?= amd64

go-build:
	go build ${BUILD_OPTS} -o bin/togo-${GOOS}-${GOARCH}

go-clean:
	go clean -x

go-deps:
	dep ensure -v

go-test:
	go test ${BUILD_OPTS} -cover ./...

docker-build:
	docker build -T ${IMAGE_NAME} .

docker-upload:
	docker push ${IMAGE_NAME}

docker-retag:
	docker pull ${IMAGE_SRC}
	docker tag ${IMAGE_SRC} ${IMAGE_DST}
	docker push ${IMAGE_DST}

bundle-all:
	# generate hash file
	echo "path  bytes  sha256" | tee ./bin/hashes
	find ./bin/ -type f -printf '%p  %s  ' -exec sh -c "sha256sum {} | cut -d ' ' -f 1;" \; | tee -a ./bin/hashes
	# compress
	tar -cvf bin/togo-all.tgz ./bin/*
	gzip ./bin/togo-darwin-* || true
	gzip ./bin/togo-linux-* || true
	# do not gzip windows binaries

bundle-clean:
	rm -v ./bin/hashes || true
	rm -v ./bin/togo-* || true