
DOCKER_IMAGE?=krishnaiyer/http-dev-server
DOCKER_TAG?=latest
OCI_REGISTRY=registry-1.docker.io
OCI_REPO=krishnaiyer
HELM_PACKAGE_NAME?=http-dev-server-helm

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

.PHONY: helm.build

helm.build:
	@echo "Build helm chart..."
	@helm package ./helm

.PHONY: helm.sign

helm.sign:
	@echo "Sign helm package ${HELM_PACKAGE_NAME}-${OCI_TAG}.tgz"
	@cd ./util/signature && go build -o "../../signature" signature.go && cd ../..
	@./signature --package ${HELM_PACKAGE_NAME}-${OCI_TAG}.tgz --private-key ${GPG_KEY_FILE} --passphrase ${GPG_PASSPHRASE}
	@rm signature

helm.push:
	@echo "Push helm chart..."
	@helm push http-dev-server-helm-${OCI_TAG}.tgz oci://${OCI_REGISTRY}/${OCI_REPO}
