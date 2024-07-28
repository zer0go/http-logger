# http-logger

## Usage
```shell
$ go run main.go
$ curl -X POST -d "hello" http://localhost:8080/write
$ tail -f /var/log/syslog | tee | while read -r line; do curl -X POST -d "$line" http://localhost:8080/write; done
$ tail -f /var/log/syslog | tee | while read -r line; do wget --quiet --output-document=- --post-data "$line" http://localhost:8080/write; done
$ curl http://localhost:8080/readAll
```
