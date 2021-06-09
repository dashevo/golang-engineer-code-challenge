APP_CMD = "cmd/app/main.go"
P2P_CMD = "cmd/p2p/main.go"
HOSTED_CMD = "cmd/hosted/main.go"
ROOT_DIR = $(CURDIR)

GO ?= go
GOCOVER  = $(GO) tool cover

BUILD_ENVS:=GOGC=off CGO_ENABLED=0

compile/app:
	$(BUILD_ENVS) $(GO) build -o ./build/app $(APP_CMD)

compile/p2p:
	$(BUILD_ENVS) $(GO) build -o ./build/p2p $(P2P_CMD)

compile/hosted:
	$(BUILD_ENVS) $(GO) build -o ./build/hosted $(HOSTED_CMD)

docker/build:
	docker compose build

docker/run/p2p:
	docker compose run -d --rm p2p-srv

docker/run/hosted:
	docker compose run -d --rm hosted-srv

run: docker/build
	docker compose run dash-challenge-app

stop:
	docker compose down

tests:
	$(GO) test -race -short -coverprofile=coverage.out -v  ./...

tests/coverage: tests
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out
