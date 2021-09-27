FROM golang:1.17

RUN apt-get update && apt-get install -y ca-certificates git-core ssh
RUN go install github.com/beego/bee/v2@latest

WORKDIR /go/src/api

EXPOSE 8080

CMD ["bee", "run"]
