FROM golang:latest

WORKDIR $GOPATH/src/github.com/oumeniOS/go-gin-blog
ADD . $GOPATH/src/github.com/oumeniOS/go-gin-blog
RUN go install .

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]
