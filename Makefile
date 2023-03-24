GOCMD:=go
GOBUILD:=$(GOCMD) build
OUTDIR:=out
TEXT:=$(shell git log -1 --pretty=%B)
NAME:=awstools-cli
MAIN:=main.go

VERSION := $(shell git describe --tags --abbrev=0)

ifeq ($(VERSION),)
NEW_VERSION=0.0.1
else
VERSION_NO_V := $(shell echo $(VERSION) | cut -c 2-)
MAJOR := $(shell echo $(VERSION_NO_V) | cut -d. -f1)
MINOR := $(shell echo $(VERSION_NO_V) | cut -d. -f2)
PATCH := $(shell echo $(VERSION_NO_V) | cut -d. -f3)
NEW_VERSION := $(MAJOR).$(shell expr $(MINOR) + 1).$(PATH)
endif

LDFLAGS=-ldflags "-X main.Version=$(NEW_VERSION)"

.PHONY: build release

build:
	@rm -fr $(OUTDIR)
	@mkdir -p $(OUTDIR)
	@echo "Building $(NAME).rpi version $(NEW_VERSION)"
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) ${LDFLAGS} -o $(OUTDIR)/$(NAME).rpi $(MAIN)
	@echo "Building $(NAME).lin version $(NEW_VERSION)"
	GOOS=linux $(GOBUILD)  ${LDFLAGS} -o $(OUTDIR)/$(NAME).lin $(MAIN)
	@echo "Building $(NAME).mac version $(NEW_VERSION)"
	GOOS=darwin $(GOBUILD) ${LDFLAGS} -o $(OUTDIR)/$(NAME).mac $(MAIN)
	zip $(OUTDIR)/$(NAME).rpi.zip $(OUTDIR)/$(NAME).rpi 
	zip $(OUTDIR)/$(NAME).lin.zip $(OUTDIR)/$(NAME).lin
	zip $(OUTDIR)/$(NAME).mac.zip $(OUTDIR)/$(NAME).mac

release:
	git tag -a v$(NEW_VERSION) -m "v$(NEW_VERSION)"	
	git push origin v$(NEW_VERSION)
	gh release create v$(NEW_VERSION) out/*.zip --tile "v$(NEW_VERSION)" --generate-notes