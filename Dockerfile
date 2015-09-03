FROM busybox
ADD ./telem_linux-amd64 /app
CMD ["/app"]
