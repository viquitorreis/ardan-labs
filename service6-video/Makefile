# checando se podemos usar ash, em imagens Alpine, ou default para BASH
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# =========================
# VAIAVEIS
KIND 			:= kindest/node:v1.29.2
KIND_CLUSTER    := ardan-starter-cluster

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

.PHONY: run
run:
	go run apis/services/sales/main.go | go run apis/tooling/logfmt/main.go