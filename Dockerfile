FROM golang:1.21.4

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy