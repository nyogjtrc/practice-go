FROM golang:1.10 as builder
COPY . /go/src/github.com/nyogjtrc/practice-go/gin/
WORKDIR $GOPATH/src/github.com/nyogjtrc/practice-go/gin/

RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -a -installsuffix cgo -o /go/bin/gin

FROM scratch
COPY --from=builder /go/bin/gin /go/bin/gin
ENTRYPOINT ["/go/bin/gin"]
EXPOSE 8080
