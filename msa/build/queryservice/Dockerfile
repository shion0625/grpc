FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN cd cmd/queryservice && go build -o /queryservice

FROM alpine:latest
COPY --from=builder /queryservice /queryservice
ENTRYPOINT ["/queryservice"]