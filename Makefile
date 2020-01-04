.PHONY: dev build image release dmg test deps clean

CGO_ENABLED=0
COMMIT=`git rev-parse --short HEAD`
APP=gopherclient
REPO?=prologic/$(APP)
TAG?=latest
BUILD?=-dev

all: dev

dev: build
	@./$(APP)

deps:
	@go get github.com/GeertJohan/go.rice/rice
	@go get ./...
	@rice embed-go

build: clean deps
	@echo " -> Building $(TAG)$(BUILD)"
	@go build -tags "netgo static_build" -installsuffix netgo \
		-ldflags "-w -X github.com/$(REPO).GitCommit=$(COMMIT) -X github.com/$(REPO).Build=$(BUILD)" .
	@echo "Built $$(./$(APP) -v)"

install:
	@go install

image:
	@docker build --build-arg TAG=$(TAG) --build-arg BUILD=$(BUILD) -t $(REPO):$(TAG) .
	@echo "Image created: $(REPO):$(TAG)"

release: dmg
	@./tools/release.sh

dmg: build
	@appify -name "GopherClient" -icon ./icon.png ./gopherclient
	@create-dmg ./GopherClient.app || true

test:
	@go test -v -cover -race $(TEST_ARGS)

clean:
	@git clean -f -x -d
