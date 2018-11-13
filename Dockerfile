FROM golang:latest

WORKDIR $GOPATH/src/github.com/oumeniOS/go-gin-blog
ADD . $GOPATH/src/github.com/oumeniOS/go-gin-blog
RUN go build .

EXPOSE 8001
ENTRYPOINT ["./go-gin-blog"]

