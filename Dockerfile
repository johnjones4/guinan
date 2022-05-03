FROM golang:1.17

WORKDIR /usr/src/app

COPY . .

RUN go get .

RUN go build .

CMD ["/usr/src/app/main"]
