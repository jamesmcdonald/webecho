FROM debian:8

RUN apt-get update && apt-get install -y ca-certificates

# Set up MySQL cloud proxy
RUN mkdir -p /opt/cloud_sql_proxy
ADD https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 /opt/cloud_sql_proxy/cloud_sql_proxy
RUN chmod 755 /opt/cloud_sql_proxy/cloud_sql_proxy
COPY credentials.json /opt/cloud_sql_proxy/credentials.json

COPY webtest /opt/webtest/webtest

COPY run-webtest.sh /run-webtest.sh

ENTRYPOINT ["/run-webtest.sh"]
