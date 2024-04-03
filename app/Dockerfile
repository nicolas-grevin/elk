ARG ELK_VERSION=7.17.18

FROM golang:1.22-alpine AS build

WORKDIR /build

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM docker.elastic.co/beats/filebeat:$ELK_VERSION

COPY --from=build /build/app /usr/local/bin/generate-logs
