# http-logger

## Usage
```shell
$ go run main.go
$ curl -X POST -d "hello" http://localhost:8080/write
$ tail -f /var/log/syslog | xargs -I %s -- curl -X POST -d "%s" http://localhost:8080/write
$ tail -f /var/log/syslog | xargs -I %s -- wget -O- --post-data "%s" http://logger:8080/write 2>&1 >/dev/null
$ curl http://localhost:8080/readAll
```
