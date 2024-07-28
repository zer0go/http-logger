# http-logger

## Usage
```shell
$ go run main.go
$ curl -X POST -d "hello" http://localhost:8080/write
$ tail -f /var/log/syslog | xargs -I %s -- curl -X POST -d "%s" http://localhost:8080/write
$ tail -f /var/log/syslog | xargs -I %s -- wget -q -O- --post-data "%s" http://localhost:8080/write
$ curl http://localhost:8080/readAll
```
