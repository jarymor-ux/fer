# GNU Make: macOS, Linux, Windows (Git Bash, MSYS2, WSL, или отдельный GNU Make).
# В cmd.exe без Unix-утилит этот Makefile может не подойти — используйте Git Bash или WSL.

GOPATH ?= $(shell go env GOPATH)
GOBIN_PATH := $(shell go env GOBIN)
GOEXE := $(shell go env GOEXE)
ifeq ($(strip $(GOBIN_PATH)),)
STATICCHECK ?= $(GOPATH)/bin/staticcheck$(GOEXE)
else
STATICCHECK ?= $(GOBIN_PATH)/staticcheck$(GOEXE)
endif

.PHONY: pre-commit fix fmt vet staticcheck staticcheck-install test check

pre-commit:
	pre-commit run --all-files -c configs/.pre-commit-config.yaml

fix:
	go fix ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck@latest

staticcheck-install:
	go install honnef.co/go/tools/cmd/staticcheck@latest

staticcheck: $(STATICCHECK)
	"$(STATICCHECK)" ./...

test:
	go test -v ./...

check: fix fmt vet staticcheck test
