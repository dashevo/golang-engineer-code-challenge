APP_CMD = "cmd/app/main.go"
ROOT_DIR = $(CURDIR)

GO ?= go

BUILD_ENVS:=GOGC=off CGO_ENABLED=0

compile:
	$(BUILD_ENVS) $(GO) build -o ./build/app $(APP_CMD)

run: compile
	./build/app
