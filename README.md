This repo consists of two major pieces: 

# Server
A server gives a json response containing the time that the response was received.

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
./bin/client -rps 1 -host "http://localhost:8080" -duration 10s

# output
```
:~/go/src/github.com/smerck/time/bin$ ./client -rps 2 -host "http://localhost:9001" -duration 10s
2020/07/27 05:28:53 Starting test: Sending 2 rps to http://localhost:9001 for 10s
2020/07/27 05:28:54 Request succeeded with Status Code 200 after 1.835278ms
2020/07/27 05:28:54 Request succeeded with Status Code 200 after 1.057445ms
2020/07/27 05:28:55 Request succeeded with Status Code 200 after 958.26µs
2020/07/27 05:28:55 Request succeeded with Status Code 200 after 1.00448ms
2020/07/27 05:28:56 Request succeeded with Status Code 200 after 1.230874ms
2020/07/27 05:28:56 Request succeeded with Status Code 200 after 1.337869ms
2020/07/27 05:28:57 Request succeeded with Status Code 200 after 1.400059ms
2020/07/27 05:28:57 Request succeeded with Status Code 200 after 1.281953ms
2020/07/27 05:28:58 Request succeeded with Status Code 200 after 1.261642ms
2020/07/27 05:28:58 Request succeeded with Status Code 200 after 1.428634ms
2020/07/27 05:28:59 Request succeeded with Status Code 200 after 1.289821ms
2020/07/27 05:28:59 Request succeeded with Status Code 200 after 1.279114ms
2020/07/27 05:29:00 Request succeeded with Status Code 200 after 891.544µs
2020/07/27 05:29:00 Request succeeded with Status Code 200 after 843.502µs
2020/07/27 05:29:01 Request succeeded with Status Code 200 after 1.343763ms
2020/07/27 05:29:01 Request succeeded with Status Code 200 after 1.275267ms
2020/07/27 05:29:02 Request succeeded with Status Code 200 after 1.279391ms
2020/07/27 05:29:02 Request succeeded with Status Code 200 after 1.295661ms
2020/07/27 05:29:03 Request succeeded with Status Code 200 after 1.314247ms
2020/07/27 05:29:03 Request succeeded with Status Code 200 after 1.303084ms
Test complete
2020/07/27 05:29:03 Overall Success Rate: 100.00%
```
