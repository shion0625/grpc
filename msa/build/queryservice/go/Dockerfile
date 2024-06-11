FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@latest

RUN cd cmd/queryservice && go build -o /queryservice

FROM alpine:latest as release
COPY --from=builder /queryservice /queryservice
COPY --from=builder /go/bin/* /usr/bin/
COPY build/queryservice/.air.toml /app/.air.toml
ENTRYPOINT ["/queryservice"]