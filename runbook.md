# Runbook: Executing Blue/Green Deploy

## Build docker image
`docker build ./server/ -t time-server`

## Publish as two images in GCR
```
docker tag time-server gcr.io/snowball-284203/time-server:v1
docker tag time-server gcr.io/snowball-284203/time-server:v2
docker push gcr.io/snowball-284203/time-server:v1
docker push gcr.io/snowball-284203/time-server:v2
```

## Apply v1 deployment and service
```
kubectl apply -f deployment-blue.yaml
kubectl apply -f service-blue.yaml
```

## Start test client
`./bin/client -rps 5 -host "http://<ip>:9001" -duration 120`

## Apply v2 deployment
`kubectl apply -f server/deployment-green.yaml`

## Apply service change to point to v2
`kubectl apply -f server/service-green.yaml`

## Spindown v1 deployment
`kubectl delete -f deployment time-server-v1`
