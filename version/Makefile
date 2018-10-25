# version
#

version := $(shell git describe --tags 2> /dev/null || echo "Unknown")
buildtime := $(shell date --utc --rfc-3339 ns| sed -e 's/ /T/')
commit := $(shell git rev-parse --short HEAD || echo "NotFound")

go_ldflags := "\
	-X github.com/nyogjtrc/practice-go/version/version.Version=$(version)\
	-X github.com/nyogjtrc/practice-go/version/version.BuildTime=$(buildtime) \
	-X github.com/nyogjtrc/practice-go/version/version.GitCommit=$(commit)"

.PHONY: build

build:
	go build -a -v -ldflags $(go_ldflags) -o ver
	go build -a -v -ldflags $(go_ldflags) -o ver-cli ./cli

run:
	go run -v -ldflags $(go_ldflags) ./*.go

protoc:
	protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto

clean:
	rm -f ./ver ./ver-cli

docker-build:
	docker build -t ver \
		--build-arg commit=$(commit) \
		--build-arg buildtime=$(buildtime) \
		--build-arg version=$(version) \
		.

docker-run:
	docker run --rm -p 8888:8888 ver
