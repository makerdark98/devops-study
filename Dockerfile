FROM alpine:3.15.4
COPY main /main
ENTRYPOINT ["/main"]