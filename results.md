

# build two docker images
sudo docker build . -t smerck/time

# publish as two images in gcr
sudo docker tag time-server gcr.io/snowball-284203/time-server:v1
sudo docker tag time-server gcr.io/snowball-284203/time-server:v2
sudo docker push gcr.io/snowball-284203/time-server:v1
sudo docker push gcr.io/snowball-284203/time-server:v2

# ran blue deployment

# ran bluegreen deployment

# ran green deployment
