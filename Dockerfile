FROM alpine:3.3
COPY queue-conf.sample.yaml /etc/seckilling/queue-conf.yaml
COPY dist/omega-queue /
COPY src/public  /public
ENTRYPOINT ["/omega-queue"]
