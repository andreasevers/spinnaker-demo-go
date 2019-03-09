FROM alpine

COPY gopath/bin/hello /hello
COPY *.jpg ./hello/
COPY *.html ./hello/

ENTRYPOINT ["/hello"]