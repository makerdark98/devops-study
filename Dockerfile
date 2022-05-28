FROM alpine:3.15.4
COPY bin/main /main
ENTRYPOINT ["/main"]
