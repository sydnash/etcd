FROM golang
ADD . /go/src/github.com/ozonru/etcd
ADD cmd/vendor /go/src/github.com/ozonru/etcd/vendor
RUN go install github.com/ozonru/etcd
EXPOSE 2379 2380
ENTRYPOINT ["etcd"]
