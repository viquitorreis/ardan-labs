# checando se podemos usar ash, em imagens Alpine, ou default para BASH
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

.PHONY: run
run:
	go run apis/services/sales/main.go | go run apis/tooling/logfmt/main.go