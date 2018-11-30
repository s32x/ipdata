FROM alpine:latest
RUN apk add --no-cache ca-certificates
ADD bin/ipdata /usr/local/bin/ipdata
CMD ipdata