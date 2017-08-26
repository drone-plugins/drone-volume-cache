# GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o drone-volume-cache
# docker build --rm -t plugins/volume-cache .
FROM scratch

ADD release/linux/amd64/drone-volume-cache /bin/
ENTRYPOINT [ "/bin/drone-volume-cache" ]
