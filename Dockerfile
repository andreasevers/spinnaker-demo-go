FROM alpine

COPY gopath/bin/hello *.jpg *.html /hello/

ENTRYPOINT ["/hello"]