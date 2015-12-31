FROM alpine:3.3
COPY octopus-conf.sample.yaml /etc/octopus/octopus-conf.yaml
COPY dist/omega-octopus /
ENTRYPOINT ["/omega-octopus"]
