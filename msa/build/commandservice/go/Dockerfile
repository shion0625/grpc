FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@latest \
    && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
    && go install github.com/volatiletech/sqlboiler/v4@latest \
    && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
    
RUN cd cmd/commandservice && go build -o /commandservice

FROM alpine:latest as release
WORKDIR /app
COPY --from=builder /commandservice /commandservice
COPY --from=builder /go/bin/* /usr/bin/
COPY build/commandservice/.air.toml /app/.air.toml
COPY internal/app/commandservice/infra/sqlboiler/config/database.toml /app/.sqlboiler.toml
ENTRYPOINT ["/commandservice"]