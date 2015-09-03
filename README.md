Telem
---

Telemetry framework

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
go build -o telem_linux-amd64 -tags netgo *.go
# Or
./script/build
```

Running
---

```bash
./telem_linux-amd64
docker run --rm -ti pdxjohnny/telem
```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen
