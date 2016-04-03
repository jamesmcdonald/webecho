#!/bin/bash

# The proxy and webtest need the following environment variables:
# TODO check these and abort if they're not set

# GCLOUD_PROJECT
#   The ID of the google cloud project

# GCLOUD_LOCATION
#   The location of the sql instance

# GCLOUD_SQL
#   The name of the cloud SQL instance

# MYSQL_USER
#   The MySQL username

# MYSQL_PASS
#   The MySQL password

# MYSQL_DB
#   The MySQL database

# URLS
#   The URLs to monitor in the format <url>;<frequency>::<url>;<frequency>

# Start the Google Cloud SQL proxy
mkdir -p /opt/cloud_sql_proxy/run
/opt/cloud_sql_proxy/cloud_sql_proxy -dir=/opt/cloud_sql_proxy/run -instances=${GCLOUD_PROJECT}:${GCLOUD_LOCATION}:${GCLOUD_SQL} -credential_file=/opt/cloud_sql_proxy/credentials.json &

# Give the proxy a moment to start up
sleep 5

# Run the webtest
/opt/webtest/webtest
