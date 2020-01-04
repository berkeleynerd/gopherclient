FROM golang:alpine AS build

RUN apk add --no-cache -U git make build-base webkit2gtk-dev

WORKDIR /src
COPY . /src

RUN make build
