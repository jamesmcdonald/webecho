BINARY=webtest
TAG_VERSION=1

all: docker

$(BINARY):
	go build

docker: $(BINARY)
	docker build -t eu.gcr.io/$(GCLOUD_PROJECT)/$(BINARY):$(TAG_VERSION) .

rm:
	rm -f $(BINARY)

rmi:
	docker rmi eu.gcr.io/$(GCLOUD_PROJECT)/$(BINARY):1

clean: rm rmi
