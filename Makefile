VERSION ?= $$(git describe --tags)

setup:
	dep ensure
	go install ./vendor/github.com/jteeuwen/go-bindata/go-bindata

build:
	go-bindata templates/
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/brewery

release: build
	tar czvf brewery.tar.gz bin
