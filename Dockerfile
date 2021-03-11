FROM golang:1.16-alpine as build

WORKDIR ./src

COPY . ./

RUN go build -mod=vendor -o=./bin/scalc . && \
    cp ./bin/scalc /usr/local/bin/ && \
    rm -rf /go/src

FROM alpine

COPY --from=build /usr/local/bin/ /usr/local/bin/

ENTRYPOINT ["scalc"]
