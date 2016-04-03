BINARY=echoserver
TAG_VERSION=1

all: docker

$(BINARY): echoserver.go
	go build

docker: $(BINARY) Dockerfile
	docker build -t eu.gcr.io/$(GCLOUD_PROJECT)/$(BINARY):$(TAG_VERSION) .

rm:
	rm -f $(BINARY)

rmi:
	docker rmi eu.gcr.io/$(GCLOUD_PROJECT)/$(BINARY):$(TAG_VERSION)

clean: rm rmi
