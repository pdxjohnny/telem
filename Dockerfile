FROM busybox
ADD ./telem_linux-386 /app
CMD ["/app"]
