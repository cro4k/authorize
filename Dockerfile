FROM golang:1.20-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache tzdata
RUN apk add git

WORKDIR /opt/app

ADD . .

ENV GOPROXY=https://goproxy.cn,direct

RUN go build -ldflags "-X github.com/cro4k/authorize/version.BuildTime=$(date '+%Y-%m-%d_%H:%M:%S') -X github.com/cro4k/authorize/version.BuildHash=$(git rev-parse HEAD)" -o authorize main.go

FROM alpine:latest AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

ENV TZ Asia/Shanghai

WORKDIR /opt/app

COPY --from=builder /opt/app/example .
COPY --from=builder /opt/app/etc/conf.yaml etc/conf.yaml

ENTRYPOINT ./example