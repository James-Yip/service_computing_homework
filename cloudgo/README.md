[Negroni]: https://github.com/urfave/negroni

[Mux]: http://www.gorillatoolkit.org/pkg/mux

# cloudgo
cloudgo is a simple web service program that return a string (with the format "hello {user}!") based on the request url.

## Framework
Since the cloudgo is really simple, I choose a tiny scheme with some lightweight components.

Scheme : `gorilla/mux` + `urfave/negroni`

- [Negroni][] is an idiomatic approach to web middleware in Go.
- [Mux][] (stands for "HTTP request multiplexer") implements a request router and dispatcher.

## Curl Test
server:
```
$ go run main.go -p 9090
[negroni] listening on :9090
```

client:
```
$ curl -v localhost:9090/James
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /James HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 12:43:35 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
<
Hello James!
* Connection #0 to host localhost left intact
```

## AB Test
```
$ ab -n 1000 -c 100 localhost:9090/James
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /James
Document Length:        13 bytes

Concurrency Level:      100
Time taken for tests:   0.123 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      130000 bytes
HTML transferred:       13000 bytes
Requests per second:    8134.18 [#/sec] (mean)
Time per request:       12.294 [ms] (mean)
Time per request:       0.123 [ms] (mean, across all concurrent requests)
Transfer rate:          1032.66 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.1      1       4
Processing:     0   11   9.1      8      43
Waiting:        0   10   9.2      7      42
Total:          0   12   9.2      8      44

Percentage of the requests served within a certain time (ms)
  50%      8
  66%     11
  75%     15
  80%     18
  90%     29
  95%     35
  98%     37
  99%     39
 100%     44 (longest request)
```

The result shows that the server served the total 1000 requests(with concurrent connections = 100) from the client within 44ms.
