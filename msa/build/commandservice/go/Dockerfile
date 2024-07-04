FROM golang:1.22.5-alpine as dev
WORKDIR /app

RUN apk --no-cache add make

RUN go install github.com/air-verse/air@latest \
    && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
    && go install github.com/volatiletech/sqlboiler/v4@latest \
    && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest \
    && go install github.com/onsi/ginkgo/v2/ginkgo@latest 

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/commandservice-linux-amd64 ./cmd/commandservice

FROM alpine:latest as release
WORKDIR /app
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=dev build/commandservice /usr/local/bin/commandservice
EXPOSE 8081
ENTRYPOINT ["commandservice"]