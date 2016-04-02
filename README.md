# webecho
An echo service and a client to test it. These are simple projects for playing with deploying containers on Google Container Engine.

## How to set up
Create a Google Cloud Platform project and note the project ID [PROJECT_ID]

Build the go binary in this directory:
    go get -d github.com/jamesmcdonald/webecho
    go build github.com/jamesmcdonald/webecho

Build the docker image:
    docker build -t eu.gcd.io/<PROJECT_ID>/echo-service

At this point you can test locally with something like:
    docker run -p 8080:8080 --rm eu.gcd.io/<PROJECT_ID>/echo-service

Push the docker image to Google:
    gcloud docker push eu.gcd.io/<PROJECT_ID>/echo-service

Choose a cluster name [CLUSTER_NAME]

Set configuration and deploy the application:
    gcloud config set project <PROJECT_ID>
    gcloud container clusters create <CLUSTER_NAME> --num-nodes=1
    gcloud container clusters get-credentials <CLUSTER_NAME>
    gcloud components install kubectl
    kubectl run echo-service --image=eu.gcr.io/<PROJECT_ID>/echo-service --port 8080
    kubectl expose deployment echo-service --type="LoadBalancer" --port 80 --target-port 8080

Check the services configuration until you see an external IP:
    kubectl get service echo-service


