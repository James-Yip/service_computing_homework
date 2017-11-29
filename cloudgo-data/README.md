[Negroni]: https://github.com/urfave/negroni

[Mux]: http://www.gorillatoolkit.org/pkg/mux

[pmlpml]: https://github.com/pmlpml/golang-learning/tree/master/web/cloudgo-data

# cloudgo-data
cloudgo-data is a simple web service program that supports the basic database services.

**note**: This program is based on the source code of teacher Pan ([pmlpml][]). What I do is just use xorm to replace the original DAO level.

## Examples

#### Insert User Infomation
```
$ curl -d "username=alice&departname=2" http://localhost:8080/service/userinfo
{
  "UID": 1,
  "UserName": "alice",
  "DepartName": "2",
  "CreateAt": "2017-11-29T23:17:20.963199603+08:00"
}

$ curl -d "username=bob&departname=2" http://localhost:8080/service/userinfo
{
  "UID": 2,
  "UserName": "bob",
  "DepartName": "2",
  "CreateAt": "2017-11-29T23:17:32.746543204+08:00"
}
```

confirm in the mysql client:
```
mysql> select * from UserInfo;
+----+----------+------------+---------------------+
| id | UserName | DepartName | CreateAt            |
+----+----------+------------+---------------------+
|  1 | alice    | 2          | 2017-11-29 23:17:20 |
|  2 | bob      | 2          | 2017-11-29 23:17:32 |
+----+----------+------------+---------------------+
2 rows in set (0.00 sec)
```


#### Find User Infomation By Id

```
$ curl http://localhost:8080/service/userinfo?userid=1
{
  "UID": 1,
  "UserName": "alice",
  "DepartName": "2",
  "CreateAt": "2017-11-30T07:17:20+08:00"
}
```


#### Find All

```
$ curl http://localhost:8080/service/userinfo?userid=
[
  {
    "UID": 1,
    "UserName": "alice",
    "DepartName": "2",
    "CreateAt": "2017-11-30T07:17:20+08:00"
  },
  {
    "UID": 2,
    "UserName": "bob",
    "DepartName": "2",
    "CreateAt": "2017-11-30T07:17:32+08:00"
  }
]
```


## AB Test
```
$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=1
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
Server Port:            8080

Document Path:          /service/userinfo?userid=1
Document Length:        102 bytes

Concurrency Level:      100
Time taken for tests:   0.678 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      235000 bytes
HTML transferred:       102000 bytes
Requests per second:    1475.92 [#/sec] (mean)
Time per request:       67.754 [ms] (mean)
Time per request:       0.678 [ms] (mean, across all concurrent requests)
Transfer rate:          338.71 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.0      0       6
Processing:     1   65  56.1     55     234
Waiting:        1   65  56.1     55     234
Total:          1   66  56.6     55     237

Percentage of the requests served within a certain time (ms)
  50%     55
  66%     70
  75%     82
  80%     92
  90%    133
  95%    212
  98%    225
  99%    228
 100%    237 (longest request)
```

The result shows that the server served the total 1000 requests(with concurrent connections = 100) from the client within 237ms.
