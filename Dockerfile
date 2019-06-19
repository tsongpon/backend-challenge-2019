FROM golang:1.12.4-alpine3.9

RUN apk add --no-cache curl
RUN apk add --no-cache git

ADD . /go/src/github.com/tsongpon/backend-challenge-2019
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/tsongpon/backend-challenge-2019

RUN dep ensure
RUN go install

EXPOSE 5000

ENTRYPOINT /go/bin/backend-challenge-2019