FROM golang:1.19.0-alpine

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod ./

COPY go.sum ./
RUN go mod download

COPY main.go .
COPY auth ./auth
COPY controllers ./controllers
COPY database ./database
COPY middlewares ./middlewares
COPY models ./models

RUN go build -o /docker-auth-ms

EXPOSE 8080

CMD ["/docker-auth-ms"]