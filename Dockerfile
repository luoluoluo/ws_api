FROM scratch
VOLUME /tmp
COPY ./main /main
ENTRYPOINT ["/main"]
