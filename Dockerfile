# Base Image
# FROM golang:alpine

# WORKDIR /app

# COPY go.mod ./
# RUN go mod tidy
# COPY . .
# RUN go build -o test

# CMD [ "./test" ]
FROM ubuntu:latest

WORKDIR /app

COPY . .

RUN apt update -y
RUN apt install redis -y
RUN service redis-server start

RUN apt install golang -y
RUN go mod tidy
RUN go build -o test

CMD [ "./test" ]
