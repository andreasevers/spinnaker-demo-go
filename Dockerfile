FROM alpine

COPY gopath/bin/hello *.jpg *.html /hello/
RUN ls -lrt /hello
RUN chmod 777 -R /hello
RUN ls -lrt /hello

ENTRYPOINT ["/hello"]