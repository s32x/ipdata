FROM alpine:latest
RUN apk add --no-cache ca-certificates
ADD bin/ipdata /usr/local/bin/server
CMD server