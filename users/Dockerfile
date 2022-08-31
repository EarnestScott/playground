FROM golang:1.19.0-alpine as builder

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

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch

COPY --from=builder /app/bin/main .

CMD ["./main"]