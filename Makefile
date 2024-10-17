SHELL := /bin/bash
APP := filesserializer
GOLANG_IMG := minixxie/golang:1.21.0

.PHONY: getimg
getimg:
	nerdctl --namespace=k8s.io images | grep local/${APP} | grep dont_push || true

.PHONY: build
build:
	nerdctl --namespace=k8s.io build . -t local/${APP}:dont_push

.PHONY: golang
golang:
	nerdctl --namespace=k8s.io run --rm -it \
		-e GOPATH=/usr/local/go \
		-v "$$PWD":/usr/local/go/src/${APP} -w /usr/local/go/src/${APP} \
		${GOLANG_IMG} sh

.PHONY: gofmt
gofmt:
	nerdctl --namespace=k8s.io run --rm -t \
		-e GOPATH=/usr/local/go \
		-v "${PWD}:/usr/local/go/src/${APP}" -w "/usr/local/go/src/${APP}" \
		"${GOLANG_IMG}" gofmt -w .

.PHONY: run
run:
	nerdctl --namespace=k8s.io run --rm -it \
		local/${APP}:dont_push
