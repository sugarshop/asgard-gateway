FROM registry.digitalocean.com/francisco/golang-base:1.18 AS build
ARG ARCH="amd64"
ARG OS="linux"
ARG PROJECT="asgard-gateway"

WORKDIR $GOPATH/src/github.com
COPY ./rpc ./rpc
COPY ./$PROJECT ./$PROJECT

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/$PROJECT
RUN sh kitex.sh && go mod tidy && sh build.sh

## release
FROM alpine:3.14
ARG PROJECT="asgard-gateway"
COPY --from=build /go/src/github.com/$PROJECT/output /app

WORKDIR /app
EXPOSE 8080
CMD ["./bootstrap.sh"]
