FROM golang:1.20.3-alpine as build-stage

WORKDIR /aggregator
COPY . .


RUN apk --no-cache update && \
apk --no-cache add git gcc libc-dev

ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
ENV GOOS=linux
ENV GOARCH=amd64

RUN export GO111MODULE=on
RUN go mod vendor
RUN go mod download 

RUN go build -tags musl -ldflags "-s -w" -o bin/aggregator ./cmd/aggregator

############################
FROM alpine
COPY --from=build-stage /aggregator/bin/aggregator .

CMD ./aggregator



