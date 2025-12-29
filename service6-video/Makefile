# checando se podemos usar ash, em imagens Alpine, ou default para BASH
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

.PHONY: run build sales dev-up dev-down dev-status dev-status-all

# =========================
# VAIAVEIS
KIND 			:= kindest/node:v1.29.2
KIND_CLUSTER    := ardan-starter-cluster
SALES_APP       := sales
VERSION         := 0.0.1
BASE_IMAGE_NAME := localhost/victorreis
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)

# =========================
# Construir containers
build: sales

sales:
	docker build \
		-f zarf/docker/dockerfile.sales \
		-t $(SALES_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

run:
	go run apis/services/sales/main.go | go run apis/tooling/logfmt/main.go