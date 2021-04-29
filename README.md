# kindle-highlight-backend


Start the server

```
go run ./main.go
```

Start auto-reload server

```
air
```

Build binary for linux

```
$ CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -o main -a -v ./main.go
```
