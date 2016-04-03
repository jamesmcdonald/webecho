# webecho
An echo service and a client to test it. These are simple projects for playing with deploying containers on Google Container Engine.

`echoserver` listens on port 8080 for HTTP requests, and replies with the URL path from the request.

`webtest` takes a list of URLs and frequencies, polls the URLs based on those frequencies and puts the resulting status and response time in a MySQL database.

## Building and testing
* If you haven't already, install the [Google Cloud SDK](https://cloud.google.com/sdk/) and run `gcloud init` to set it up.

* Create a Google Cloud Platform project and note the project ID. Stick it in `$GCLOUD_PROJECT` and configure gcloud to use it.
```
    export GCLOUD_PROJECT=echo-service-12345
    gcloud config set project $GCLOUD_PROJECT
```

**Note: You can build containers and test them without creating a project, but you should set `$GCLOUD_PROJECT` to something because the Makefiles use it to tag the Docker images.**

* The Docker build for `webtest` requires a file with service account credentials called `credentials.json`. Create a service account selecting JSON credentials. It should get 'editor' permissions by default, which is what is necessary. Copy the downloaded credentials file into `webtest/credentials.json`.

**Note: These credentials give access to your Google Cloud project. Be careful where you put them and the generated Docker images that will contain them. There might be a smarter way to handle this but I haven't found it yet :D**

* Run `make`. This will build the binaries and create Docker images.

* At this point you can test locally with something like:
```
    docker run -p 8080:8080 --rm eu.gcr.io/$GCLOUD_PROJECT/echoserver:1
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
    kubectl run echo-service --image=eu.gcr.io/$GCLOUD_PROJECT/echoserver:1 --port 8080
    kubectl expose deployment echo-service --type="LoadBalancer" --port 80 --target-port 8080
```

* Check the services configuration until you see an external IP. Once you see it, you should be able to point a browser there.
```
    kubectl get service echo-service
```

## Deploying `webtest` to Google

* Create a Cloud SQL instance. It must be a second generation instance to support the Cloud SQL proxy. Permit access for your own IP to connect to the database, and set the root password.

* Connect to MySQL and create a database and user.
```
    mysql -u root -h <sql_cloud_ip> -p
    create database echostats;
    grant all privileges on echostats.* to echostats@'%' identified by '<mysql_password>';
```

* For this simple example there isn't any clever schema management, so just create the table too.
```
    use echostats;
    create table echostats (
        timestamp datetime,
        url varchar(200),
        host varchar(100),
        status integer,
        duration integer,
        primary key (timestamp, url, host)
    );
```

* Again, you can test locally with something like this.
```
    docker run -ti --rm --env URLS="http://10.11.12.13;300" \
                        --env MYSQL_HOST=<sql_cloud_ip> \
                        --env MYSQL_USER=echostats \
                        --env MYSQL_PASSWORD=<mysql_password> \
                        --env MYSQL_DB=echostats \
                        eu.gcr.io/echo-service-1267/webtest:2
```

The `URLS` environment variable should be set to pairs of a URL and a frequency to test it in seconds. URLs should be separated by `::`. For example
```
    URLS="http://10.11.12.13/;300::http://10.11.12.14/;600"
```
This will start a cloud sql proxy with incorrect configuration, but it won't be used. You should see entries appear in the database with a `hostname` that matches the id of the running container.

* Push the docker image to Google.
```
    gcloud docker push eu.gcr.io/$GCLOUD_PROJECT/webtest
```

* Create a deployment to run the container. To monitor the echoserver, use its URL in the URLS parameter. Change any parameters you need.
```
    kubectl run webtest --image=eu.gcr.io/$GCLOUD_PROJECT/webtest:2 \
                        --env URLS="https://10.11.12.12;27::http://10.11.12.13/;99" \
                        --env GCLOUD_PROJECT=$GCLOUD_PROJECT \
                        --env GCLOUD_LOCATION=europe-west1 \
                        --env GCLOUD_SQL=echostats \
                        --env MYSQL_USER=echostats \
                        --env MYSQL_DB=echostats \
                        --env MYSQL_PASSWORD=<mysql_password>
```

You should see entries start to appear in the MySQL database after a few seconds. The `hostname` will be the name of the Kubernetes pod available in `kubectl get pods`. You can scale up the deployment, or deploy it to other locations to compare the response times. Or deploy different configurations with different frequencies for different URLs!
