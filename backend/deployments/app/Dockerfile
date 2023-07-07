FROM golang:1.20.5-bullseye

WORKDIR /app

COPY ./go.mod ./

RUN apt-get update \
    && apt-get install git \
    && go get github.com/labstack/echo/v4 \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/zerolog/log \
    && go install github.com/cosmtrek/air@latest \