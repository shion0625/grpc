FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN cd cmd/commandservice && go build -o /commandservice

FROM alpine:latest
COPY --from=builder /commandservice /commandservice
ENTRYPOINT ["/commandservice"]