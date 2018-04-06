FROM alpine:3.6
EXPOSE 53/udp
ADD ./tmpdns /
ENTRYPOINT ["/tmpdns"]
