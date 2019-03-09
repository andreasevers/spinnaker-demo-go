FROM alpine

COPY gopath/bin/hello /hello
COPY . /hello

ENTRYPOINT ["/hello"]