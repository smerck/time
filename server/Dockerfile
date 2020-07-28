FROM golang:alpine 

WORKDIR $GOPATH/src/smerck/time/server
COPY . .

RUN apk add --no-cache git mercurial \
    && go get -v github.com/gorilla/mux \
    && apk del git mercurial
RUN go build -o ./bin/server
RUN go install 

EXPOSE 9001
ENTRYPOINT ["/go/bin/server"]