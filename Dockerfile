# GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o drone-local-cache
# docker build --rm -t plugins/local-cache .
FROM scratch

ADD drone-local-cache /bin/
ENTRYPOINT [ "/bin/drone-local-cache" ]
