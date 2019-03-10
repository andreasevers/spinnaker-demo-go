FROM alpine

COPY gopath/bin/hello *.jpg *.html /hello/
RUN useradd -m -g hadoop hadoop
RUN echo 'hadoop:hadoop' | chpasswd
RUN chown hadoop:hadoop -R /hello/

ENTRYPOINT ["/hello"]