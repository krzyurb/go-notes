FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o /build/main .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["/build/main"]
