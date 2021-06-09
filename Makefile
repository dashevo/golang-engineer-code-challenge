APP_CMD = "cmd/app/main.go"
P2P_CMD = "cmd/p2p/main.go"
SELF_HOSTED_CMD = "cmd/selfhosted/main.go"
ROOT_DIR = $(CURDIR)

GO ?= go
GOCOVER  = $(GO) tool cover

BUILD_ENVS:=GOGC=off CGO_ENABLED=0

compile/app:
	$(BUILD_ENVS) $(GO) build -o ./build/app $(APP_CMD)

compile/p2p:
	$(BUILD_ENVS) $(GO) build -o ./build/p2p $(P2P_CMD)

compile/self-hosted:
	$(BUILD_ENVS) $(GO) build -o ./build/self_hosted $(SELF_HOSTED_CMD)

docker/build:
	docker compose build

docker/run/p2p:
	docker compose run -d --rm p2p-srv

docker/run/self-hosted:
	docker compose run -d --rm self-hosted-srv

run: docker/build
	docker compose run dash-challenge-app

stop:
	docker compose down

tests:
	$(GO) test -race -short -coverprofile=coverage.out -v  ./...

tests/coverage: tests
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

lint:
	golangci-lint run ./...
