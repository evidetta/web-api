FROM alpine:3.8
RUN apk add --no-cache libc6-compat

COPY ./bin/web-api /usr/bin
COPY ./bin/migrations /migrations

ENTRYPOINT /usr/bin/web-api
