FROM golang:alpine

RUN echo $GOPATH

WORKDIR $GOPATH/src/smerck/time/server
COPY . .

RUN go get -d -v
RUN go build -o ./bin/server
RUN go install 

EXPOSE 9001
ENTRYPOINT ["/go/bin/server"]