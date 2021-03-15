FROM golang:1.16-alpine as build

WORKDIR ./src

COPY . ./

RUN cp -r ./test /usr/local/bin/test

RUN go build -mod=vendor -o=./bin/scalc ./cmd/... && \
    cp ./bin/scalc /usr/local/bin/ && \
    rm -rf /go/src

FROM alpine

COPY --from=build /usr/local/bin/ /usr/local/bin/
COPY --from=build /usr/local/bin/test /usr/local/bin/test

WORKDIR /usr/local/bin/test

ENTRYPOINT ["scalc"]
