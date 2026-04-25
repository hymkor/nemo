ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=set
    WHICH=where.exe
    DEL=del
    NUL=nul
    CP=cmd /c copy
else
    SET=export
    WHICH=which
    DEL=rm
    NUL=/dev/null
    CP=cp
endif

ifndef GO
    SUPPORTGO=go1.20.14
    GO:=$(shell $(WHICH) $(SUPPORTGO) 2>$(NUL) || echo go)
endif

NAME:=$(notdir $(CURDIR))
VERSION:=$(shell git describe --tags 2>$(NUL) || echo v0.0.0)
GOOPT:=-ldflags "-s -w -X github.com/hymkor/$(NAME).version=$(VERSION)"
EXE:=$(shell $(GO) env GOEXE)

build:
	$(GO) fmt ./...
	$(SET) "CGO_ENABLED=0" && $(GO) build $(GOOPT) -tags debug "./cmd/$(NAME)"

all:
	$(GO) fmt ./...
	$(SET) "CGO_ENABLED=0" && $(GO) build $(GOOPT) "./cmd/$(NAME)"

test:
	$(GO) fmt ./...
	$(GO) test -v ./...

_dist:
	$(SET) "CGO_ENABLED=0" && $(GO) build $(GOOPT) "./cmd/$(NAME)"
	zip -9 $(NAME)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(NAME)$(EXE)

dist:
	$(SET) "GOOS=windows" && $(SET) "GOARCH=386"   && $(MAKE) _dist
	$(SET) "GOOS=windows" && $(SET) "GOARCH=amd64" && $(MAKE) _dist
	$(SET) "GOOS=darwin"  && $(SET) "GOARCH=amd64" && $(MAKE) _dist
	$(SET) "GOOS=darwin"  && $(SET) "GOARCH=arm64" && $(MAKE) _dist
	$(SET) "GOOS=freebsd" && $(SET) "GOARCH=amd64" && $(MAKE) _dist
	$(SET) "GOOS=linux"   && $(SET) "GOARCH=386"   && $(MAKE) _dist
	$(SET) "GOOS=linux"   && $(SET) "GOARCH=amd64" && $(MAKE) _dist

bump:
	$(GO) run github.com/hymkor/latest-notes@latest -suffix "-goinstall" -gosrc $(NAME) CHANGELOG*.md > version.go

clean:
	$(DEL) *.zip $(NAME)$(EXE)

release:
	$(GO) run github.com/hymkor/latest-notes@latest | gh release create -d --notes-file - -t $(VERSION) $(VERSION) $(wildcard $(NAME)-$(VERSION)-*.zip)

manifest:
	$(GO) run github.com/hymkor/make-scoop-manifest@latest -all *-windows-*.zip > $(NAME).json

.PHONY: all test dist _dist clean release manifest docs draft
