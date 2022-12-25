FROM golang:1.19-bullseye

RUN apt update && apt-get -y install iproute2

RUN mkdir /opt/fpm-logger
ADD go.mod main.go /opt/fpm-logger/

WORKDIR /opt/fpm-logger
RUN go build && mv fpm-logger /usr/local/bin
