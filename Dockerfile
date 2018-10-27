FROM alpine:3.8
RUN apk add --no-cache libc6-compat
COPY ./bin/web-api /usr/bin
ENTRYPOINT /usr/bin/web-api
