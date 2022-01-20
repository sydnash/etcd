FROM golang
ADD . /go/src/github.com/sydnash/etcd
ADD cmd/vendor /go/src/github.com/sydnash/etcd/vendor
RUN go install github.com/sydnash/etcd
EXPOSE 2379 2380
ENTRYPOINT ["etcd"]
