This repo consists of two major pieces: 

# Server
A server gives a json response containing the time that the request was received.

# build
go build -o ./bin/server ./server/server.go

# test
go test ./...

## Client
A client to test the server. The client takes a few arguments:

* rps: (Default: 10) Number of requests to send to the server per second.
* host: (Default: localhost:9001) host to target with requests. Should start with the protocol (http[s]://)
* duration: (Default: 120s)

The client will send requests to the specified server and indicate if the request succeeded/failed and how long it took to complete.
Example output:

# build
go build -o ./bin/client ./client/client.go

# run 
./bin/client -rps 1 -host "http://localhost:8080" -duration 10

# output
```
go run client.go -host http://localhost:8080 -rps 100 -duration 120
2020/07/27 09:13:10 Starting test: Sending 100 requests per second to http://localhost:8080 for 120s.
2020/07/27 09:13:10 Request 1 succeeded with Status Code 200 after 896.773µs
2020/07/27 09:13:10 Request 2 succeeded with Status Code 200 after 518.112µs
2020/07/27 09:13:10 Request 3 succeeded with Status Code 200 after 491.168µs
[...]
2020/07/27 08:53:53 Request 11998 succeeded with Status Code 200 after 369.933µs
2020/07/27 08:53:53 Request 11999 succeeded with Status Code 200 after 353.273µs
2020/07/27 08:53:53 Request 12000 succeeded with Status Code 200 after 348.487µs
2020/07/27 08:53:53 Test complete
2020/07/27 08:53:53 Overall Success Rate: 100.00%
2020/07/27 08:53:53 Total Number of Requests: 12000
```
