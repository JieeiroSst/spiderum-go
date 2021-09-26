# Build stage
FROM golang:1.16-alpine3.13 AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/http/main.go
RUN apk --no-cache add curl

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY . .
COPY --from=builder /app/main .
COPY start.sh .
COPY wait-for.sh .

EXPOSE 5004
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]