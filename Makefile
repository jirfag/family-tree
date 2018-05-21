.PHONY: all example test

export PROJECT_PATH = $GOPATH/src/family-tree

all: test

example:
	export GIN_MODE=test && go test -v -cover ./handler

test: example
	go test -v -cover ./handler

docker_test: clean
	docker run --rm \
		-v $(PWD):$(PROJECT_PATH) \
		-w=$(PROJECT_PATH) \
		fredliang/golang-testing \
		sh -c "coverage all"

clean:
	rm -rf .cover