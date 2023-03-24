GOCMD:=go
GOBUILD:=$(GOCMD) build
OUTDIR:=out
VERSION:=$(shell git ls-remote --refs --sort="version:refname" --tags | cut -d/ -f3- | tail -n1 | xargs -I{} echo '{} + 1' | bc)
TEXT:=$(shell git log -1 --pretty=%B)
NAME:=awstools-cli
MAIN:=main.go

ifeq ($(VERSION),)
VERSION=v0.0.1
endif

LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

build:
	@rm -fr $(OUTDIR)
	@mkdir -p $(OUTDIR)
	@echo "Building $(NAME).rpi version $(VERSION)"
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) ${LDFLAGS} -o $(OUTDIR)/$(NAME).rpi $(MAIN)
	@echo "Building $(NAME).lin version $(VERSION)"
	GOOS=linux $(GOBUILD)  ${LDFLAGS} -o $(OUTDIR)/$(NAME).lin $(MAIN)
	@echo "Building $(NAME).mac version $(VERSION)"
	GOOS=darwin $(GOBUILD) ${LDFLAGS} -o $(OUTDIR)/$(NAME).mac $(MAIN)
	zip $(OUTDIR)/$(NAME).rpi.zip $(OUTDIR)/$(NAME).rpi 
	zip $(OUTDIR)/$(NAME).lin.zip $(OUTDIR)/$(NAME).lin
	zip $(OUTDIR)/$(NAME).mac.zip $(OUTDIR)/$(NAME).mac

release:
	git tag -a $(VERSION) -m "v$(VERSION)"
	git push origin $(VERSION)
	hub release create -a $(OUTDIR)/$(NAME).rpi.zip -a $(OUTDIR)/$(NAME).lin.zip -a $(OUTDIR)/$(NAME).mac.zip -m "$(TEXT)" $(VERSION)