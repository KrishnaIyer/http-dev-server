
DOCKER_IMAGE?=krishnaiyer/http-dev-server
DOCKER_TAG?=latest

.PHONY: init

init:
	@echo "Initialize..."
	@go mod download && go mod tidy

.PHONY: docker.build

docker.build:
	@echo "Build docker image..."
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker.push

docker.push:
	@echo "Push docker image..."
	@docker push $(DOCKER_IMAGE):$(DOCKER_TAG)
