FROM alpine

COPY gopath/bin/hello ./hello/
COPY *.jpg ./hello/
COPY *.html ./hello/
RUN chmod a+x -R /hello
RUN ["chmod", "+x", "/hello"]
RUN ["chmod", "+x", "/hello/hello"]

ENTRYPOINT ["chmod 755 /hello && /hello"]