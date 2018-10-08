export PROJECT_PATH = $GOPATH/src/family-tree

all: test

env:
	export GIN_MODE=test && \
	swag init

test: env
	go test -v -cover ./handler

cov: test
	go test -race ./handler -coverprofile=coverage.txt -covermode=atomic

deploy: env
	GOOS=linux GOARCH=amd64  go build -tags=jsoniter ./main.go && \
	docker build -t $(DOCKER_REGISTRY)/fredliang/family-tree  .  && \
	docker push $(DOCKER_REGISTRY)/fredliang/family-tree

clean:
	rm -rf coverage.txt main docs