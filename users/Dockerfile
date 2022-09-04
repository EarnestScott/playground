# start from golang base image
FROM golang:1.19.0-alpine as builder

# enable go modules
ENV GO111MODULE=on
ENV DB_CONN=localhost
ENV PORT=8080

# set current working directory
WORKDIR /app

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.
COPY go.mod ./
COPY go.sum ./

# download all dependencies
RUN go mod download

# copy source code
COPY . .

# CG-_ENABLES is disabled for cross system compilation
# it is also common best practice
# build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# start a new stage from scratch
FROM scratch

# copy pre-built binary file
COPY --from=builder /app/bin/main .

#run executable
CMD ["./main", "-conn $DB_CONN", "-port $PORT"]