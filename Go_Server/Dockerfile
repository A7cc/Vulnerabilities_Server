FROM golang:1.21.1-alpine

RUN apk add --no-cache git bash

WORKDIR /app

COPY src/ .

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn

RUN go mod download

EXPOSE 8081

CMD ["go", "run", "."]
