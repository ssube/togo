GOOS ?= linux
GOARCH ?= amd64

go-build:
	go build ${BUILD_OPTS} -o bin/togo-${GOOS}-${GOARCH}

go-clean:
	go clean -x

go-test: go-build
	go test ${BUILD_OPTS} -cover -coverprofile=out/cover.out ./...
	go tool cover -html=out/cover.out -o=out/cover.html
	go tool cover -func=out/cover.out

git-push:
	git push github
	git push gitlab

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
	find ./bin/ -name "togo-*" -type f -printf '%p  %s  ' -exec sh -c 'sha256sum $$1 | cut -d " " -f 1;' find-exec {} \; \
		| tee -a ./bin/hashes
	# compress
	tar -cvf bin/togo-all.tgz ./bin/*
	gzip ./bin/togo-darwin-* || true
	gzip ./bin/togo-linux-* || true
	# do not gzip windows binaries

bundle-clean:
	rm -v ./bin/hashes || true
	rm -v ./bin/togo-* || true

upload-climate:
	cc-test-reporter after-build \
		--debug \
		-r "$(shell echo "${CODECLIMATE_SECRET}" | base64 -d)" \
		-t gocov

upload-codecov:
	codecov --disable=gcov \
		--file=out/cover.out \
		--token=$(shell echo "${CODECOV_SECRET}" | base64 -d)
