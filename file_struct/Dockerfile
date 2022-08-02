FROM golang:latest

RUN mkdir "/app"

WORKDIR /app

ADD . /app

RUN go build -o server main.go

EXPOSE 9999
CMD ./server
