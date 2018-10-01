FROM alpine:3.7
RUN apk add --no-cache ca-certificates
WORKDIR /

ADD prometheus_server /prometheus/
ADD prometheus.yml /etc/prometheus/prometheus.yml

EXPOSE 9090

CMD ["/prometheus/prometheus", "--config.file", "/etc/prometheus/prometheus.yml"]