FROM golang:1.16.2-alpine as build

MAINTAINER luckyziv

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR /app/go-shopping-demo

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/ .

EXPOSE 3002

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

CMD ["./out"]
