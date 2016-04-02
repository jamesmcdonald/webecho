FROM debian:8

RUN apt-get update && apt-get install -y ca-certificates

COPY webtest /opt/webtest/webtest

ENTRYPOINT ["/opt/webtest/webtest"]
