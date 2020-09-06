FROM golang:1.13 as builder

COPY . /go/src/github.com/nadavb/gobulenat

WORKDIR /go/src/github.com/nadavb/gobulenat/api/server

ENV CGO_ENABLED=0
ENV GO111MODULE=off
ENV GOOD=linux

RUN go build -o /go/bin/gobulenat

FROM alpine:3.7

COPY --from=builder /go/bin/gobulenat /app/gobulenat

WORKDIR /app

ADD api/server/templates templates/
ADD api/server/static static/

CMD /app/gobulenat