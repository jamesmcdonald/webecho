# webecho
An echo service and a client to test it. These are simple projects for playing with deploying containers on Google Container Engine.

## Building and testing
* If you haven't already, install the [Google Cloud SDK](https://cloud.google.com/sdk/) and run `gcloud init` to set it up.

* Create a Google Cloud Platform project and note the project ID. Stick it in `$GCLOUD_PROJECT` and configure gcloud to use it.
```
export GCLOUD_PROJECT=echo-service-12345
gcloud config set project $GCLOUD_PROJECT
```

Note: You can build containers and test them without creating a project, but you should set `$GCLOUD_PROJECT` to something because the Makefiles use it to tag the Docker images.

* Run `make`. This will build the binaries and create Docker images.

* At this point you can test locally with something like:
```
    docker run -p 8080:8080 --rm eu.gcr.io/$GCLOUD_PROJECT/echoserver
```

## Deploying `echoserver` to Google
* Push the docker image to Google.
```
    gcloud docker push eu.gcr.io/$GCLOUD_PROJECT/echoserver
```

* Choose a cluster name. Set it as `$GCLOUD_CLUSTER`.
```
export GCLOUD_CLUSTER=echo-service
```

* Create a container cluster and get its credentials.
```
    gcloud container clusters create $GCLOUD_CLUSTER --num-nodes=1
    gcloud container clusters get-credentials $GCLOUD_CLUSTER
```

* Install the Kubernetes management tool and fire up the containers. I call the deployment `echo-service` here. You can use something else if you prefer.
```
    gcloud components install kubectl
    kubectl run echo-service --image=eu.gcr.io/$GCLOUD_PROJECT/echoserver --port 8080
    kubectl expose deployment echo-service --type="LoadBalancer" --port 80 --target-port 8080
```

* Check the services configuration until you see an external IP. Once you see it, you should be able to point a browser there.
```
    kubectl get service echo-service
```
