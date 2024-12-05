FROM golang:1.23.3
WORKDIR /usr/src/app

RUN groupadd -g 1009 go-api
RUN useradd -g go-api -m go-api

COPY . .

RUN chown -R go-api:go-api ./

USER go-api:go-api

RUN go install github.com/air-verse/air@latest
RUN go install github.com/dgraph-io/badger/v4/badger@latest
RUN go mod tidy