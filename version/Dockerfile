FROM golang:1.11 as builder
COPY . /go/src/github.com/nyogjtrc/practice-go/version/
WORKDIR $GOPATH/src/github.com/nyogjtrc/practice-go/version/

RUN go get
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ARG version=na
ARG buildtime=na
ARG commit=na
RUN go build -a -v -ldflags "\
    -X github.com/nyogjtrc/practice-go/version/version.Version=${version}\
    -X github.com/nyogjtrc/practice-go/version/version.BuildTime=${buildtime}\
    -X github.com/nyogjtrc/practice-go/version/version.GitCommit=${commit}\
" -o /go/bin/ver

RUN go build -a -v -ldflags "\
    -X github.com/nyogjtrc/practice-go/version/version.Version=${version}\
    -X github.com/nyogjtrc/practice-go/version/version.BuildTime=${buildtime}\
    -X github.com/nyogjtrc/practice-go/version/version.GitCommit=${commit}\
" -o /go/bin/ver-cli ./cli

FROM scratch
COPY --from=builder /go/bin/ver /ver
COPY --from=builder /go/bin/ver-cli /ver-cli
ENTRYPOINT ["/ver"]
EXPOSE 8888
