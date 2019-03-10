FROM alpine

COPY gopath/bin/hello ./hello/
COPY *.jpg ./hello/
COPY *.html ./hello/
RUN chmod a+x -R /hello

ENTRYPOINT ["/hello"]