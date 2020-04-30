PACKAGE = github.com/floydeconomy/arisaedo-go

GIT_COMMIT = $(shell git --no-pager log --pretty="%h" -n 1)
GIT_TAG = $(shell git tag -l --points-at HEAD)
ARISAEDO_VERSION = $(shell cat cmd/arisaedo/VERSION)

PACKAGES = `go list ./... | grep -v '/vendor/'`

MAJOR = $(shell go version | cut -d' ' -f3 | cut -b 3- | cut -d. -f1)
MINOR = $(shell go version | cut -d' ' -f3 | cut -b 3- | cut -d. -f2)
export GO111MODULE=on

.PHONY: arisaedo all clean test

arisaedo:| go_version_check
	@echo "building $@..."
	@go build -v -o $(CURDIR)/bin/$@ -ldflags "-X main.version=$(ARISAEDO_VERSION) -X main.gitCommit=$(GIT_COMMIT) -X main.gitTag=$(GIT_TAG)" ./cmd/arisaedo
	@echo "done. executable created at 'bin/$@'"

dep:| go_version_check
	@go mod download

go_version_check:
	@if test $(MAJOR) -lt 1; then \
		echo "Go 1.13 or higher required"; \
		exit 1; \
	else \
		if test $(MAJOR) -eq 1 -a $(MINOR) -lt 13; then \
			echo "Go 1.13 or higher required"; \
			exit 1; \
		fi \
	fi

all: arisaedo

clean:
	-rm -rf \
$(CURDIR)/bin/arisaedo \

test:| go_version_check
	@go test -cover $(PACKAGES)

