FROM golang:alpine

RUN mkdir /app

ENV CGO_ENABLED 0

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o /build/main .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["/build/main"]
