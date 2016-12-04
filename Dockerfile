FROM alpine
MAINTAINER jspc <james@zero-internet.org.uk>

ADD go-json-echo-http-server /
ENTRYPOINT ["/go-json-echo-http-server"]
