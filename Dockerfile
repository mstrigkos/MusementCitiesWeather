# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY config.json ./
RUN go build -o /MusementCitiesWeather

CMD [ "/MusementCitiesWeather" ]