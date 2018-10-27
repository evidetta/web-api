REPO_PREFIX=minikube
SERVICE_NAME=web-api
VERSION=latest

BINDIR=./bin

GC=go build
GC_OPTS=-o $(BINDIR)/$(SERVICE_NAME)
CC=docker build
CC_OPTS=-t $(REPO_PREFIX)/$(SERVICE_NAME):$(VERSION) ./

all: compile containerize

compile:
	mkdir -p $(BINDIR)
	$(GC) $(GC_OPTS)

containerize:
	$(CC) $(CC_OPTS)

clean:
	rm -r $(BINDIR)/*
