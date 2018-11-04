FROM alpine:3.8
MAINTAINER Orange

RUN apk add --update --no-cache supervisor bash tzdata

RUN mkdir -p /etc/supervisor.d /logs
COPY supervisord.conf /etc
