FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

EXPOSE 5432

EXPOSE 4000

CMD go run ./cmd/web


