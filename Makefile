tag ?= latest
version ?= $(shell yq e '.version' helm/chart/Chart.yaml)
clean-cmd = docker compose down --remove-orphans --volumes

binary:
	go build -o im-database-manager -ldflags "-s -w" ./cmd/serve

check:
	pre-commit run --all-files --show-diff-on-failure

smoke-test:
	docker compose up -d database jwks
	sleep 3
	IMAGE_TAG=$(tag) docker compose up -d prod

docker-image:
	IMAGE_TAG=$(tag) docker compose build prod

init:
	direnv allow
	pip install pre-commit
	pre-commit install --install-hooks --overwrite

push-docker-image:
	IMAGE_TAG=$(tag) docker compose push prod

dev:
	docker compose up --build dev database jwks

cluster-dev:
	skaffold dev

test: clean
	docker compose up -d database jwks
	docker compose run --no-deps test
	$(clean-cmd)

dev-test: clean
	docker compose run --no-deps dev-test
	$(clean-cmd)

clean:
	$(clean-cmd)
	go clean

helm-chart:
	@helm package helm/chart

publish-helm:
	@curl --user "$(CHART_AUTH_USER):$(CHART_AUTH_PASS)" \
        -F "chart=@im-database-manager-$(version).tgz" \
        https://helm-charts.fitfit.dk/api/charts

swagger-check-install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger-clean:
	rm -rf swagger/sdk/*
	rm -f swagger/swagger.yaml

swagger-docs: swagger-check-install
	swagger generate spec -o swagger/swagger.yaml -x swagger/sdk --scan-models
	swagger validate swagger/swagger.yaml

swagger-client: swagger-check-install
	swagger generate client -f swagger/swagger.yaml -t swagger/sdk

swagger: swagger-clean swagger-docs swagger-client

.PHONY: binary check docker-image init push-docker-image dev test dev-test helm-chart publish-helm
