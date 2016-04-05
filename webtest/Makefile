BINARY=webtest
TAG_VERSION=2

all: docker

$(BINARY): webtest.go
	go get
	go build

docker: $(BINARY) Dockerfile
	docker build -t eu.gcr.io/$(GCLOUD_PROJECT)/$(BINARY):$(TAG_VERSION) .

rm:
	rm -f $(BINARY)

rmi:
	docker rmi eu.gcr.io/$(GCLOUD_PROJECT)/$(BINARY):$(TAG_VERSION)

clean: rm rmi
