FROM debian:8

#RUN apt-get update
#RUN apt-get -y dist-upgrade

COPY echoserver /opt/echoserver/echoserver

EXPOSE 8080

ENTRYPOINT ["/opt/echoserver/echoserver"]
