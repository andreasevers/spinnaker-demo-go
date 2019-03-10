FROM alpine

ADD gopath/bin /hello
# ADD *.jpg ./hello/
# ADD *.html ./hello/

ENTRYPOINT ["/hello"]