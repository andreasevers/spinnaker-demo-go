FROM alpine

COPY gopath/bin/hello *.jpg *.html /hello/
RUN ls -lrt .
RUN chmod 777 -R .
RUN ls -lrt .

ENTRYPOINT ["/hello"]