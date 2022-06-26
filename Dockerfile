FROM golang:1.16 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# EXPOSE 8080 8080
WORKDIR /go/src/application
CMD ["air"]