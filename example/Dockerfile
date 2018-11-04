FROM myalpine:0.7

RUN mkdir /services
COPY example /services
COPY example.ini /etc/supervisor.d

EXPOSE 7070

ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
