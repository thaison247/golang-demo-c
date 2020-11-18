FROM golang:1.13.4-alpine
MAINTAINER Hung Son
RUN apk add git \
    && mkdir /build
COPY . /build
RUN cd /build && go build -o main
RUN apk add --no-cache -U tzdata bash ca-certificates \
    && update-ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && chmod 711 /build/main \
    && rm -rf /var/cache/apk/*
WORKDIR /build
CMD ["/build/main"]