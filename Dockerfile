FROM alpine

COPY gopath/bin/hello /hello
COPY gopath/src/hello /hello

ENTRYPOINT ["/hello"]