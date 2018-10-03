FROM alpine:3.6

RUN apk add --no-cache nginx-mod-http-lua

# Delete default config
RUN rm -r /etc/nginx/conf.d && rm /etc/nginx/nginx.conf

# Create folder for PID file
RUN mkdir -p /run/nginx

# Add our nginx conf
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

RUN mkdir /etc/nginx/nginx-lua-prometheus
ADD ./nginx/prometheus.lua /etc/nginx/nginx-lua-prometheus/

ADD ./nginx/metrics.conf /etc/nginx/conf.d/

CMD ["nginx"]