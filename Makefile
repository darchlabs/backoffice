# load .env file
include backoffice.env
export $(shell sed 's/=.*//' backoffice.env)

SERVICE_NAME=backoffice
DOCKER_USER=darchlabs

build:
	@echo "[building backoffice]"
	@docker build -t darchlabs/backoffice -f ./Dockerfile --progress tty .
	@echo "Build darchlabs/backoffice docker image done ✔︎"

build-pristine:
	@echo "[building image backoffice]"
	@docker build --no-cache -t darchlabs/backoffice -f ./Dockerfile --progress tty .
	@echo "Build darchlabs-backoffice docker image done ✔︎"

apply-dev:
	@echo "[applying infra/dev/$(art).yaml]"
	@kubectl apply -f infra/dev/$(art).yaml
	@echo "Applied darchlabs/backoffice-dev ✔︎"

compose-up:
	@echo "[composing backoffice up]"
	@docker-compose -f docker-compose.yml up

compose-down:
	@echo "[composing backoffice down]"
	@docker-compose -f docker-compose.yml down

build-local:
	@echo "[build darchlabs/backoffice local]"
	@go build -o bin/backoffice/backoffice cmd/backoffice/main.go
	@echo "Build darchlabs-backoffice done ✔︎"

dev:
	@echo "[run backoffice local]"
	@export $$(cat backoffice.env) && nodemon --exec go run cmd/server/main.go

docker-login:
	@echo "[docker] Login to docker..."
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)

docker: docker-login
	@echo "[docker] pushing $(REGISTRY_URL)/$(SERVICE_NAME):$(VERSION)"
	@docker buildx create --use
	@docker buildx build --platform linux/amd64,linux/arm64  --push -t $(DOCKER_USER)/$(SERVICE_NAME):$(VERSION) .

create-migration:
	@echo "[create migration]"
	@goose -dir=migrations/ create $(name)
	@echo "migration migrations/$(name) created ✔︎"

test:
	@echo "[TEST]"
	@export $$(cat test.env) && go test -p 1 -failfast -cover -race -v -count=1 ./...
	@echo "Done ✔︎"


