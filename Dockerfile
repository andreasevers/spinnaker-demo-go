FROM alpine

COPY gopath/bin/hello *.jpg *.html /hello/
RUN ls -lrt /hello
RUN chmod +x /hello
RUN ls -lrt /hello

ENTRYPOINT ["/hello"]