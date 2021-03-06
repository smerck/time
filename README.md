# Goals
* Create a server that returns the current timestamp/date in json format.
* Create a client that can query this API at a specified rate, and record success, failure, and/or TTLB.
* Execute a blue-green deploy without a failed request.

# Server 
Returns json output with the server time for the received request.

## Quickstart
  ### build
  `make server`

  ### test
  `make test`

  ### run 
  `./bin/server &`

  ### build docker container
  `make docker-build`

# Client
A client to test the server. The client takes a few arguments:

* rps: (Default: 10) Number of requests to send to the server per second.
* host: (Default: localhost:9001) host to target with requests. Should start with the protocol (http[s]://)
* duration: (Default: 120s)

The client will send requests to the specified server and indicate if the request succeeded/failed and how long it took to complete. At the end of test completion, the client will log the overall success rate and the total number of requests sent.

## Quickstart
  ### build
  `make client`

  ### run
  `./bin/client  -host http://localhost:8080 -rps 10 -duration 120`

  ### example output
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

# Executing Blue/Green Deploys
For this, I set up a [GKE](https://cloud.google.com/kubernetes-engine) cluster, pushed my server container to [GCR](https://cloud.google.com/container-registry), and executed my bluegreen deployment with kubectl. I've included a runbook with my intended steps and results from running the blue/green deploy.

  * Runbook for executing blue/green deploys with this service [runbook.md](https://github.com/smerck/time/blob/master/runbook.md)
  * Results available in [results.md](https://github.com/smerck/time/blob/master/results.md)
