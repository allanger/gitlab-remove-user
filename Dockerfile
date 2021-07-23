# Building
FROM golang:1.16.6-alpine3.14 as builder
WORKDIR /go/src/app
COPY . /go/src/app
RUN apk add git gcc
RUN go build main.go


FROM  alpine:3.14
WORKDIR /root/
COPY --from=builder /go/src/app/main .
ENTRYPOINT ["./main"]  