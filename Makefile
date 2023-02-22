tag ?= latest
version ?= $(shell yq e '.version' helm/chart/Chart.yaml)
clean-cmd = docker compose down --remove-orphans --volumes

init:
	direnv allow
	pip install pre-commit
	pre-commit install --install-hooks --overwrite

check:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	pre-commit run --all-files --show-diff-on-failure

docker-image:
	IMAGE_TAG=$(tag) docker compose build prod

push-docker-image:
	IMAGE_TAG=$(tag) docker compose push prod

dev:
	docker compose up --build dev database jwks

test: clean
	docker compose run --no-deps test
	$(clean-cmd)

test-coverage: clean
	docker compose run --no-deps test-coverage
	$(clean-cmd)

smoke-test:
	docker compose up -d database jwks
	sleep 3
	IMAGE_TAG=$(tag) docker compose up -d prod

clean:
	$(clean-cmd)
	go clean

swagger-check-install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	swagger version

swagger-clean:
	rm -rf swagger/sdk/*
	rm -f swagger/swagger.yaml

swagger-docs: swagger-check-install
	swagger generate spec -o swagger/swagger.yaml -x swagger/sdk --scan-models
	swagger validate swagger/swagger.yaml

swagger-client: swagger-check-install
	swagger generate client -f swagger/swagger.yaml -t swagger/sdk
	goimports -v -w swagger/sdk/

swagger: swagger-clean swagger-docs swagger-client

.PHONY: check docker-image init push-docker-image dev test
