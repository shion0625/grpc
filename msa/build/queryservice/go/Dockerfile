# dev
FROM golang:1.22.5-alpine as dev
WORKDIR /app

RUN apk --no-cache add make

RUN go install github.com/air-verse/air@latest \
    && go install github.com/onsi/ginkgo/v2/ginkgo@latest

COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/queryservice-linux-amd64 ./cmd/queryservice

# release
FROM alpine:latest as release
WORKDIR /app
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=dev build/queryservice /usr/local/bin/queryservice
ENTRYPOINT ["queryservice"]