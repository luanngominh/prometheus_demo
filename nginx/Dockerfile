#ref: https://stackoverflow.com/questions/47366214/how-do-i-add-the-lua-module-for-nginx-on-alpine-linux
FROM alpine:3.6

RUN apk add --no-cache nginx-mod-http-lua

# Delete default config
RUN rm -r /etc/nginx/conf.d && rm /etc/nginx/nginx.conf

# Create folder for PID file
RUN mkdir -p /run/nginx

# Add our nginx conf
COPY ./nginx.conf /etc/nginx/nginx.conf

CMD ["nginx"]