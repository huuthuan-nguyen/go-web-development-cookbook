FROM golang:1.11

ENV SRC_DIR="/go/src/cookbook"
ENV GOBIN="/go/bin"

WORKDIR $GOBIN

ADD . $SRC_DIR

RUN cd /go/src;

RUN go install github.com/huuthuan-nguyen/go-web-development-cookbook

ENTRYPOINT ["./cookbook"]

EXPOSE 8080