# Executing Blue/Green Deploy

# build docker image
`sudo docker build . -t smerck/time`

# publish as two images in gcr, we're using the same docker image and time-server binary,
# but can still demonstrate a blue/green deploy with two containers.
```
sudo docker tag time-server gcr.io/snowball-284203/time-server:v1
sudo docker tag time-server gcr.io/snowball-284203/time-server:v2
sudo docker push gcr.io/snowball-284203/time-server:v1
sudo docker push gcr.io/snowball-284203/time-server:v2
```

# apply blue deployment
```
kubectl apply -f deployment-blue.yaml
kubectl apply -f service-blue.yaml
```

# start test client
`./bin/client -rps 5 -host "http://105.198.89.14:9001" -duration 120`

# apply bluegreen deployment
`kubectl apply -f deployment-bluegreen.yaml`

# apply green deployment
`kubectl apply -f service-green.yaml`
