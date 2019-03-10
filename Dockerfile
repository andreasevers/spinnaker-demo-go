FROM alpine

ADD gopath/bin/hello *.jpg *.html /hello/
# ADD *.jpg ./hello/
# ADD *.html ./hello/

ENTRYPOINT ["/hello"]