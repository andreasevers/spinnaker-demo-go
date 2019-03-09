FROM alpine

COPY gopath/bin/hello *.jpg *.html /hello/
RUN chmod +x /hello

ENTRYPOINT ["/hello"]