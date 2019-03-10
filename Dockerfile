FROM alpine

COPY gopath/bin/hello /hello
ADD *.jpg ./hello/
ADD *.html ./hello/

ENTRYPOINT ["/hello"]